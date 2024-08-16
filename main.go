package main

import (
	"fmt"
)

func main() {
    a := 10
    b := 5

    fmt.Println("Addition:", calculator.Add(a, b))
    fmt.Println("Subtraction:", calculator.Sub(a, b))
    fmt.Println("Multiplication:", calculator.Mul(a, b))
    fmt.Println("Division:", calculator.Divide(a, b))
}
