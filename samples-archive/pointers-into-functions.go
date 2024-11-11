/*
### Passing Pointers to a Function in Go
=========================================
- Pointers in Go programming language or Golang is a variable which is used to store the memory address of another variable.

- You can also pass the pointers to the function like the variables. There are two ways to do this as follows:

1. Create a pointer and simply pass it to the Function
2. Passing an address of the variable to the function call

*/


package main

import (
   "fmt"
)

func pointer_funct(a *int){
	// a function that takes in an integer type pointer as a parameter

	// change the value stored in that memory address
	*a = 666
	fmt.Println("Memory Address = ", a)  // pr &v1
}


func main(){

    fmt.Println("\n### ==> Changing value stored in a certain memory address")
	var v1 = 11
	var pnt_v = &v1
	fmt.Println("Initial Value stored in v1 = ", v1)
    fmt.Println("Address of y = ", pnt_v)  // pr &v1
	fmt.Println("Initial Value stored in pnt_v OR &v1 = ", *pnt_v)

	pointer_funct(pnt_v)  // Pass into the function a created pointer. You could also just pass the address.. i,e pointer_funct(&v1)
	fmt.Println("New v1 value after Changing = ",v1)


}