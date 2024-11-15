## Remember a Type called struct?

In Go, a struct itself does not contain methods, but methods can be associated with a struct type. 
Methods are functions that have a special receiver argument, which ties them to a particular type, like a struct.
To associate methods with a struct, you define functions with a receiver of that struct type:

package main

import "fmt"

// Define a struct
type Person struct {
    Name string
    Age  int
}

// Define a method on the Person struct
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}

// Another method that modifies the struct's Age field
func (p *Person) HaveBirthday() {
    p.Age++
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    fmt.Println(person.Greet()) // "Hello, my name is Alice"
    person.HaveBirthday()
    fmt.Println(person.Age) // 31
}

## use type struct or use type interface?
- Use a Struct when you need a concrete data type with specific fields, such as Person, Rectangle, or Circle.
  Struct defines concrete data structures with fields to hold data.

- Use an Interface when you want to define a set of behaviors (methods) that multiple types can implement, allowing for  polymorphism and flexibility in your code.
An Interface specifies a set of method signatures without implementations.

## Interfaces
- In Go, an interface is a type that is a collection of a set of methods without providing their code/implementation.
  By not providing the implementation, it allows different types to implement the same interface differently.

- By defining a set of method signatures that any type can implement, interfaces provide a flexible and powerful way to write generic code.

- Any type that implements those methods is considered to satisfy the interface, making interfaces key for achieving polymorphism and flexible code in Go.

- You can’t create an instance of an interface directly, but you can make a variable of the interface type to store any value that has the needed methods.


