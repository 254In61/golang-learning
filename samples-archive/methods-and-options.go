/*
### overview
============
=> Go methods are like functions but with a key difference: they have a receiver argument, which allows access to the receiver’s properties.
=> The receiver can be a struct or non-struct type, but both must be in the same package. 
=> Methods cannot be created for types defined in other packages, including built-in types like int or string; otherwise, the compiler will raise an error.

Syntax
func(receiver_name Type)  method_name(parameter_list) (return_type) { 
// Code
}
Receiver: Can be accessed within the method.

=> looks like when defining python class and the methods..

class XYZ():
   def __init__(self,a,b):
      self.a = a
      self.b = b

   // Define methods under this class
   def getSomething(self):
     // some statements here

*/

package main

import "fmt"

/*
Method with a struct reciever
==============================

In Go, you can define a method where the receiver is of a struct type. 
The receiver is accessible inside the method. 
The example below showcases this approach with a struct type receiver.
*/

// Defining a struct
type employee struct {
    fname, sname, title string
    age,salary int
}

// Defining a method with a struct reciever
func (emp employee) display(){
    fmt.Println("First Name : ", emp.fname)
    fmt.Println("Second Name : ", emp.sname)
    fmt.Println("Job Title : ", emp.title)
    fmt.Println("Age : ", emp.age)
    fmt.Println("Salary : ", emp.salary)
}

/*
Method with a non-struct receiver
=================================

Go also allows methods with non-struct type receivers, as long as the receiver’s type and method definition are present in the same package. 
You cannot define a method with a receiver type from another package (e.g., int, string).
*/

// Creating a custom type based on int
type number int

// Defining a method with a non-struct receiver
func (n number) getSquare() number {
    return n * n
}

/*
Methods with Pointer Receiver
=============================
In Go, methods can have pointer receivers. 
This allows changes made in the method to reflect in the caller.

func (p *Type) method_name(...Type) Type {    
    // Code
}
*/

// Defining a struct
type person struct {
    name string
}

// Method with pointer receiver to modify data
func (p *person) changeName(newName string) {
    p.name = newName
}

func main(){

    // 1) Methods with Struct Type Receiver
    
    fmt.Println("###==>Methods with Struct Type Receiver")
    e_1 := employee{ "Allan", "Maseghe", "Golang Develooper", 42, 210000 }
    // fmt.Println("employee 1 :",e_1)

    // Calling the method
    e_1.display()

    // 2) Methods with Non-Struct Type Receiver
    
    fmt.Println("\n###==>Methods with Non-Struct Type Receiver")
    a := number(4)
    b := a.getSquare()

    fmt.Println("Square of", a, "is", b)

    // 3. Method with pointer receiver to modify data
    fmt.Println("\n###==>Method with pointer receiver to modify data")
    a := person{name: "allan"}
    fmt.Println("Before:", a.name)

    // Calling the method to change the name
    a.changeName("burukenge")
    fmt.Println("After:", a.name)
}


