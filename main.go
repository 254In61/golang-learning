/*
### overview
============
=> A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly
different types into a single type.
=> Pointers in Go programming language or Golang is a variable which is used to store the memory address of another variable.

*/

package main

import "fmt"

type Employee struct {
    fname, sname, title string
    age,salary int
}

func main(){

    // usual pointer usage
    x := 333
    pnt_x := &x
    fmt.Println("memory address :",&x)
    fmt.Println("pointer value :",pnt_x)

    // declaring and initializing a struct using a struct literal
    e_1 := Employee{ "Allan", "Maseghe", "Golang Develooper", 42, 210000 }
    fmt.Println("employee 1 :",e_1)

    // what if we want to change struct field values within the memory address
    // You can also create a pointer to a struct as shown: 
    // pnt_emp1 is a pointer to the Employee struct
    pnt_emp66 := &Employee{ "David", "Maungao", "K8s specialist", 45, 250000 } 
    fmt.Println("pointer value :",pnt_emp66)

    // Access struct fields using the pointer.
    fmt.Println("First Name:", (*pnt_emp66).fname)
    fmt.Println("Age:", (*pnt_emp66).age)
    fmt.Println("Pay:", (*pnt_emp66).salary)


}
