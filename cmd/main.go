// cmd/main.go

package main

import (
    "fmt"
    "os"
    "strconv"
    "arthicli/pkg/calculator"
)

func main() {
    args := os.Args[1:]
    if len(args) != 3 {
        fmt.Println("Usage: arithcli <operation> <num1> <num2>")
        os.Exit(1)
    }

    operation, num1Str, num2Str := args[0], args[1], args[2]
    num1, err1 := strconv.Atoi(num1Str)
    num2, err2 := strconv.Atoi(num2Str)
    if err1 != nil || err2 != nil {
        fmt.Println("Please provide valid numbers")
        os.Exit(1)
    }

    switch operation {
    case "add":
        result := calculator.Add(num1, num2)
        fmt.Println("Result:", result)
    case "subtract":
        result := calculator.Subtract(num1, num2)
        fmt.Println("Result:", result)
    case "multiply":
        result := calculator.Multiply(num1, num2)
        fmt.Println("Result:", result)
    case "divide":
        result, err := calculator.Divide(num1, num2)
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(1)
        }
        fmt.Println("Result:", result)
    default:
        fmt.Println("Unknown operation:", operation)
        os.Exit(1)
    }
}
