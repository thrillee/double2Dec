Certainly, here's a README in Markdown format for your Go code:

# SQL Double to Decimal Conversion Tool

This Go program provides a simple way to generate SQL queries for converting double/float columns to decimal columns with a specified precision.

## Usage

1. **Installation**: You need to have Go installed. If not, you can download and install it from [https://golang.org/dl/](https://golang.org/dl/).

2. **Clone the Repository**: Clone this repository to your local machine.

   ```bash
   git clone https://github.com/your-username/sql-double-to-decimal.git
   cd sql-double-to-decimal
   ```

3. **Build the Program**: Build the Go program using the following command:

   ```bash
   go build
   ```

4. **Run the Program**: Run the program with the required options.

   ```bash
   ./sql-double-to-decimal -t [TABLE] -c [COLUMN1 COLUMN2 ...] -d [DECIMAL PROPERTY] -nn
   ```

   - `-t`, `--table`: Name of the table containing the columns you want to convert.
   - `-c`, `--columns`: Space-separated names of the columns you want to convert.
   - `-d`, `--decimal`: (Optional) The decimal property to apply (e.g., "11,2"). If not specified, the default is "(11,2)".
   - `-nn`, `--notnull`: (Optional) Add this flag if the column is not null.

5. **Example**:

   ```bash
   ./sql-double-to-decimal -t my_table -c price discount -d "(10,2)" -nn
   ```

   This will generate SQL queries to convert the "price" and "discount" columns in the "my_table" table to decimal(10,2) format, assuming they are not null.

## Author

- Oluwatobi Bello(https://github.com/thrillee)

## License

This project is open-source and available under the [MIT License](LICENSE). Feel free to use, modify, and distribute it as needed.

---

Feel free to replace `[TABLE]`, `[COLUMN1 COLUMN2 ...]`, `[DECIMAL PROPERTY]`, and `[Your Name]` with your actual values and information.

You can also add more details, such as examples, in-depth explanations, or advanced use cases to your README as needed.
