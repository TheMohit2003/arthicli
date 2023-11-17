---
![Alt text](public/golang.png)
# arthicli: Arithmetic CLI Tool

## Introduction

`arthicli` is a command-line interface (CLI) tool built in Go, designed to perform basic arithmetic operations. It serves as a practical example of Go's capabilities in CLI development, including package management, unit testing, and modular architecture.

## Features

- Perform arithmetic operations such as addition, subtraction, multiplication, and division.
- Simple and intuitive command-line interaction.
- Robust error handling for invalid inputs and edge cases.
- Scalable project structure, allowing for further expansion and inclusion of additional features.

## Getting Started

### Prerequisites

Ensure you have Go installed on your system. This project requires Go version 1.16 or newer. You can check your Go version using:

```bash
go version
```

### Installation

Clone the `arthicli` repository to your local machine (assuming you have the project hosted on a platform like GitHub):

```bash
git clone https://github.com/yourusername/arthicli.git
cd arthicli
```

Build the project using:

```bash
go build -o arthicli ./cmd
```

This will create an executable named `arthicli`.

### Usage

Run arithmetic operations by executing the `arthicli` command followed by the operation and operands. For example:

```bash
./arthicli add 5 3       # Output: Result: 8
./arthicli multiply 6 4  # Output: Result: 24
```

Available operations are `add`, `subtract`, `multiply`, and `divide`.

## Running the Tests

To ensure the functionality of `arthicli` is working as expected, you can run the unit tests:

```bash
go test ./...
```

This command will execute all tests in the project.

---
