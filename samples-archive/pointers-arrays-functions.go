/*
### Passing arrays into functions
=========================================
=> When you pass an array directly to a function, Go creates a full copy of the array. 
Any changes made to this array inside the function do not affect the original array.

=> This behavior can lead to higher memory usage and slower performance, especially with large arrays. 

** Remember : Arrays are avoided since they are fixed length. But we will show about arrays then we go to slices
=> Passing arrays by value can be memory-intensive for large arrays, as it creates a duplicate in memory.

*/


package main

import (
   "fmt"
)

func modifyArray_normal(arr_x [5]int) {
	// Here, array is passed in as the parameter.
	// Code just duplicates the original array and changes it, but the original one is NOT changed
    arr_x[1] = 99
	arr_x[3] = 666
}

func modifyArray_with_pntr(arr_x *[5]int){
    // Function takes in a pointer
	// Data is changed on the memory level
    (*arr_x)[1] = 99
	(*arr_x)[3] = 666
	
}



func main() {
    myArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println("\n###==> Pass array as complete data to a function... Original isn't changed")
    modifyArray_normal(myArray)
    fmt.Println(myArray) // Output: [1 2 3 4 5], original array still remains unchanged
    
	fmt.Println("\n###==> Pass pointer to the function... Original is changed")
    modifyArray_with_pntr(&myArray)
    fmt.Println(myArray) // Output: [1 2 3 4 5], original array still remains unchanged
}
