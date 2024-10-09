# Domain Processing Utility

This Go program processes domain names to extract their effective top-level domain plus one (eTLD+1) using the [publicsuffix](https://github.com/globalsign/publicsuffix) library.

## Prerequisites

- Go 1.21 or later

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/yourusername/domain-processing-utility.git
   cd domain-processing-utility
   ```

2. Initialize the Go module:
   ```
   go mod init domain-processing-utility
   ```

3. Create a `go.mod` file with the following content:
   ```
   module domain-processing-utility

   go 1.21

   require github.com/globalsign/publicsuffix v0.0.0-20180223084127-7115ee24d247
   ```

4. Install dependencies:
   ```
   go mod tidy
   ```

## Building the Program

To build the program, run:

```
go build -o apexparser
```

This will create an executable named `apexparser` in your current directory.

## Usage

The program reads domain names from standard input (stdin), processes them, and prints the results to standard output (stdout).

### Interactive Mode

Run the program and type domain names, one per line:

```
./apexparser
example.com
www.example.co.uk
```

Press Ctrl+D (Unix/Linux/Mac) or Ctrl+Z (Windows) to exit.

### Using echo

Process a single domain:

```
echo "www.example.com" | ./apexparser
```

Process multiple domains:

```
echo -e "www.example.com\nexample.co.uk" | ./apexparser
```

### Reading from a File

Process domains stored in a file:

```
cat domains.txt | ./apexparser
```

Or:

```
./apexparser < domains.txt
```

### Saving Output to a File

To save the processed results to a file:

```
./apexparser < domains.txt > results.txt
```

## Error Handling

- The program will print warnings for invalid inputs or processing errors to stderr.
- It will continue processing subsequent inputs even if an error occurs.
- Fatal errors will cause the program to exit with a non-zero status code.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

