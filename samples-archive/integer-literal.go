package main

import "fmt"

func main() {
    // Decimal literal
    var decimal int = 42
    fmt.Println("Decimal:", decimal) // Output: 42

    // Binary literal
    var binary int = 0b1010
    fmt.Println("Binary:", binary) // Output: 10

    // Octal literal
    var octal int = 075
    fmt.Println("Octal:", octal) // Output: 61

    // Hexadecimal literal
    var hex int = 0x2A
    fmt.Println("Hexadecimal:", hex) // Output: 42
}

/*

Integer Literals and Type Inference
Go automatically infers the integer type if you use shorthand syntax or if the literal is used in a context that requires a specific integer type:

go
Copy code
// Default to int type
num := 123        // Inferred as int

// Specified type
var bigNum int64 = 1234567890123456789
Underscores for Readability
Starting with Go 1.13, you can use underscores (_) in integer literals to make large numbers more readable:

go
Copy code
var largeNumber = 1_000_000  // Same as 1000000
fmt.Println(largeNumber)     // Output: 1000000
Summary
Integer literals in Go are flexible and allow for different bases to suit your needs. By supporting binary, octal, decimal, and hexadecimal literals, Go provides a straightforward way to represent values in the form most suitable for the context of the code.








*/