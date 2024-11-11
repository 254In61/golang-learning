/*
=> functions

In Go, a function is declared with the func keyword, followed by its name, parameters, and optional return type.

Syntax
func function_name(Parameter-list)(Return_type) {
    // function body...
}
For our multiply example:

func multiply(a, b int) int {
    return a * b
}

*/


package main

import (
   "fmt"
)

func multiply_normal(a, b int) int {
	a = a * 2 // modifying a inside the function
    return a * b
}

func multiply_pntr(a, b *int) int {
    *a = *a * 2 // modifying a's value at its memory address
    return *a * *b
}


func main(){
    
	// array_name:= [length]Type{item1, item2, item3,...itemN}
	var y,z int
	num_slice := []int{7,22,88,100,76,276}
	z = 5

	/*
    ## Call function by Value
    - In call by value, values of the arguments are copied to the function parameters, 
	  so changes in the function do not affect the original variables.

	*/

	// Iterate array using for loop 
	fmt.Println("####=> Call function by Value")
	for i:=0; i < len(num_slice); i++{ 
        fmt.Printf("%d\n", multiply_normal(z , num_slice[i])) 
    } 

	/*
    #### Call by Reference
    - In call by reference, pointers are used so that changes inside the function reflect in the caller’s 
	 variables.
	*/
    
    fmt.Println("####=> Call function by Reference")
	for i:=0; i < len(num_slice); i++{
		y = num_slice[i]
        fmt.Printf("Before: z = %d, y = %d\n", z, y) 
		result := multiply_pntr(&z, &y)
		fmt.Printf("multiplication: %d\n", result)
        fmt.Printf("After: z = %d, y = %d\n", z, y)
    } 
    
	

}