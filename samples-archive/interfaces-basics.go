/*
## Interfaces
==============
- In Go, an interface is a type that is a collection of a set of methods without providing their code/implementation.
  By not providing the implementation, it allows different types to implement the same interface differently.

- By defining a set of method signatures that any type can implement, interfaces provide a flexible and powerful way to 
  write generic code.

- Any type that implements those methods is considered to satisfy the interface, making interfaces key for 
  achieving polymorphism and flexible code in Go.

- You can’t create an instance of an interface directly, but you can make a variable of the interface 
  type to store any value that has the needed methods.

*/

package main

import (
  "fmt"
  "math"
)

// define the interface
type Shape interface{
    // You define a set of methods.
    // These methods will then be implemented down the line. 
    // Remember : A method is a function with a reciever. The reciever will be of any Type of your choice.
    Area() float64
    Perimeter() float64
}

// #### struct type that implements the Shape interface.

/*
### STRUCT

=> A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly
different types into a single type.
=> Any real-world entity which has some set of properties/fields can be represented as a struct. 
=> This concept is generally compared with the classes in object-oriented programming.
   It can be termed as a lightweight class that does not support inheritance but supports composition. 

=> For Example, an address has a name, street, city, state, Pincode. 
   It makes sense to group these three properties into a single structure address as shown below.

    type Address struct {
      Name string 
      Street string
      Code int
    }

=> You interract with the struct like Classes in Python.. Using .
   add_1 =: Address{Name: "Maseghe", Steet: "8 isabella st", Code: 4301}
   
    ==> accessing struct fields")
    fmt.Println("Addr1 Name     : ", add_1.Name)
    fmt.Println("Addr1 Street   : ", add_1.Age)
    fmt.Println("Addr1 Code     : ", add_1.Code)

### METHODS and STRUCTS

=>  Methods are functions that have a special receiver argument, which ties them to a particular type, like a struct.
=>  In Go, a struct itself does not contain methods, but methods can be associated with a struct type. 
    To associate methods with a struct, you define functions with a receiver of that struct type
*/


/* 
## Circle type that implements the Shape interface

1. Define a struct type Circle
2. define a method whose reciever is the struct type and the method implements the Shape interface
*/

// define a struct type Circle
// This struct will serve as the reciever input for the methods(listed under Interface) as they are implemented.

type Circle struct{
    radius float64
}



/*
Method with a struct reciever
==============================

In Go, you can define a method where the receiver is of a struct type. 
The receiver is accessible inside the method. 
The example below showcases this approach with a struct type receiver.

Implement the methods collection within the Interface.
Remember a method is a function with a reciever. This time, the reciever is the struct above.

*/

func (c Circle) Area() float64 { 
    return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.radius
}


func main(){

    var s_1 Shape
    s_1 = Circle{radius: 5}
    fmt.Println("C Area:", s_1.Area())
    fmt.Println("C Perimeter:", s_1.Perimeter())

}

