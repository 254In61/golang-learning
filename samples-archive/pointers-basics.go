/*
Pointers in Go programming language or Golang is a variable which is used to store the memory address of another variable.

### Change value directly in an address
=========================================
- You can change the value that is stored within a memory address/location instead of assigning a new value to the variable.
    
=> When passing data to a function, one makes a copy of the data ( check pointers-into-functions.md )
If you have large structs, slices, or arrays, duplicating them means large overheads and inefficient memory utilizations.

Being able to change the data directly in the memory, saves the need for duplication of data.
If using functions, one just passes the pointer to the function for the function to do the rest of the changes.
       
Advantages : 
	   1. Efficient Memory Usage: For example, passing a large struct as a pointer to a function avoids duplicating it, reducing memory usage.

	   2. Performance Gains: Modifying data in place through pointers can be faster than making copies. 
	      This is especially true in performance-critical applications,such as systems programming or handling network requests, 
		  where high throughput and low latency are important.

	   3. Function Flexibility and Efficiency: In Go, functions can modify the values they receive through pointers. 
	      This means you can update a variable directly in memory rather than relying on a return value or a more complicated workaround. 
		  It’s efficient and often leads to simpler code.
		
	   4. Enhanced Control Over Data:Go’s pointers allow you to decide exactly when a variable should be copied versus when it should be modified in place. 
          This level of control can lead to optimized, predictable behavior.

When to Use Direct Memory Access (Pointers) in Go
=====================================================

1. For Large or Complex Data Types:

If you have large structs, slices, or arrays, passing them as pointers is generally more efficient. 
You avoid the overhead of duplicating the data, which is especially relevant if the data size is significant.

2. Mutating Functions:

When you write functions that need to change the value of an argument, passing a pointer is the way to go. 
This lets the function modify the original variable, making the code cleaner and often more intuitive.

3. Concurrency and Shared State:

When you have shared data structures across goroutines, using pointers can help control access to the data. 
However, this should be combined with synchronization primitives (like mutexes) to ensure thread safety.


Caution When Using Direct Memory Access
=========================================
While direct memory manipulation can lead to optimized code, it also introduces risks:

1. Safety and Readability: 

Go’s explicit pointer syntax is designed to make it clear when you’re working with memory addresses. 
Excessive use of pointers can make code harder to understand, especially for developers who may be unfamiliar with pointer semantics.


2. Memory Leaks and Bugs: 

Improper use of pointers can lead to memory leaks or bugs due to unintended modifications of shared data.

3. Garbage Collection Overhead: 

Although Go has garbage collection, excessive pointer use can sometimes add complexity to garbage collection, as Go tracks memory references.

*/


package main

import (
   "fmt"
)


func main(){

    fmt.Println("\n### ==> Changing value stored in a certain memory address")
	var v1 = 777
	var pnt_v = &v1

	fmt.Println("Initial Value stored in v1 = ", v1)
    fmt.Println("Address of y = ", pnt_v)  // pr &v1
	fmt.Println("Initial Value stored in pnt_v OR &v1 = ", *pnt_v)

	// change the value by assigning the new value to the pointer.
	*pnt_v = 111
	fmt.Println("New v1 value after Changing = ",v1)

	

}