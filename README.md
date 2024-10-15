# JAKTVA (JSON Add Key To Value All)

JAKTVA is a Go program that adds a specified key-value pair to all objects in a JSON file, including nested objects.

## Features

- Add a key-value pair to all objects in a JSON file
- Support for nested objects and arrays
- Automatic type inference for values (boolean, number, string)
- Optional datatype specification for strict type checking

## Download

You can download the latest pre-compiled binaries for your platform:

- [Linux](releases/0.0.1/jaktva-0.0.1-linux)
  - [macOS](releases/0.0.1/jaktva-0.0.1-osx)
- [Windows](releases/0.0.1/jaktva-0.0.1-win.exe)

## Usage

### Parameters:

- `-i`: Input JSON file path (required)
- `-k`: Key to add to each object (required)
- `-v`: Value to assign to the key (required)
- `-t`: Optional datatype (bool, int, float, string)

### Examples:

1. Add a string value:
   ```
   go run main.go -i input.json -k "newKey" -v "newValue"
   ```

2. Add a number with automatic type inference:
   ```
   go run main.go -i input.json -k "count" -v "42"
   ```

3. Add a boolean with specified datatype:
   ```
   go run main.go -i input.json -k "isActive" -v "true" -t bool
   ```

## Output

The program creates a new JSON file with "-modified" appended to the original filename, containing the updated JSON structure.

## Building

To build the program for multiple platforms, use the provided `build.sh` script:

```
This will create executables for Linux, macOS, and Windows in the `releases/<version>` directory.

