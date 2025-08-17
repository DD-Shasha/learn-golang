# Golang Learning Workspace

This repository contains various Go (Golang) learning projects and examples, including basic programs and a starter CRUD app.

## Prerequisites

- Go 1.25 or later installed ([Download Go](https://golang.org/dl/))
- (Optional) [Gorilla Mux](https://github.com/gorilla/mux) for REST API examples

## Project Structure

- `100_Ques/` - Basic Go exercises
- `stephen/` - Step-by-step Go learning modules
- `crud-app/` - Starter CRUD REST API example

## How to Run Go Projects

### Method 1: Directly Run

```sh
go run main.go [other_go_files.go]
```

### Method 2: Build and Run

1. Initialize the module (if not already done):
   ```sh
   go mod init <module-name>
   ```
2. Download dependencies:
   ```sh
   go mod tidy
   ```
3. Build the project:
   ```sh
   go build -o <output-binary>
   # Example:
   go build -o hello
   ```
4. Run the compiled binary:
   ```sh
   ./hello
   ```

## Example: Running the CRUD App

```sh
cd crud-app
go mod tidy
go build -o crud-app
./crud-app
```

The CRUD app will start a server on port 8000. You can test the endpoints using curl or Postman.

---

Happy Coding!
