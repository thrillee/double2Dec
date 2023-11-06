package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type NewColumn struct {
	addColumnQuery string
	tmpColumn      string
	tableName      string
	ogColumnName   string
	decimalProp    string
	nullableParam  string
}

func createTmpColQuery(tableName, columnName, decimalProp string, notNull bool) NewColumn {
	nullableParam := "null"
	if notNull {
		nullableParam = "not null"
	}

	tmpColumn := fmt.Sprintf("%s_dec", columnName)

	addQuery := fmt.Sprintf(
		"alter table %s add %s %s %s after %s;", tableName, tmpColumn, decimalProp, nullableParam, columnName)

	return NewColumn{
		addColumnQuery: addQuery,
		ogColumnName:   columnName,
		tmpColumn:      tmpColumn,
		tableName:      tableName,
		decimalProp:    decimalProp,
		nullableParam:  nullableParam,
	}
}

func moveRecords(newColumn NewColumn) string {
	return fmt.Sprintf(
		"update %s set %s=%s where %s > 0;",
		newColumn.tableName, newColumn.tmpColumn, newColumn.ogColumnName, newColumn.ogColumnName)
}

func migrateToOldColumn(newColumn NewColumn) string {
	dropQuery := fmt.Sprintf(
		"alter table %s drop column %s;", newColumn.tableName, newColumn.ogColumnName)

	changeColumnQuery := fmt.Sprintf(
		"alter table %s change %s %s %s %s;",
		newColumn.tableName, newColumn.tmpColumn,
		newColumn.ogColumnName, newColumn.decimalProp, newColumn.nullableParam)

	return fmt.Sprintf("%s\n%s", dropQuery, changeColumnQuery)
}

func main() {
	app := &cli.App{
		Name:  "SQL Double to Decimal",
		Usage: "This program help create sql query that you can use to convert a double/float column to decimal(11,2)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "table",
				Aliases:  []string{"t"},
				Usage:    "Table",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     "columns",
				Aliases:  []string{"c"},
				Usage:    "Columns names",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "decimal",
				Aliases:  []string{"d"},
				Usage:    "Decimal Property e.g (11,2)",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "notnull",
				Aliases:  []string{"nn"},
				Usage:    "Column is not null",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			decimalValue := "decimal"
			decimalProp := ctx.String("decimal")
			if decimalProp == "" {
				decimalProp = "(11,2)"
			}
			descimalQuery := decimalValue + decimalProp

			columns := ctx.StringSlice("columns")
			table := ctx.String("table")

			isNotNull := ctx.Bool("notnull")

			for _, column := range columns {
				if column == "" {
					continue
				}

				fmt.Printf("-- COLUMN: %s \n", column)

				newColumn := createTmpColQuery(table, column, descimalQuery, isNotNull)
				fmt.Printf("-- Tmp Column: %s\n", newColumn.tmpColumn)

				fmt.Println(newColumn.addColumnQuery)
				fmt.Println(moveRecords(newColumn))
				fmt.Println(migrateToOldColumn(newColumn))
				fmt.Println("")
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
