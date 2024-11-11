/*
=> When you pass an array directly to a function, Go creates a full copy of the array. 
Any changes made to this array inside the function do not affect the original array.
This behavior can lead to higher memory usage and slower performance, especially with large arrays. 

=> Remember : Arrays are avoided since they are fixed length. If creating a function, the function
parameters are fixed to that array size, making it not scalable.

func modifyArray_normal(arr_x [5]int) {
	// Here, array is passed in as the parameter.
	// Code just duplicates the original array and changes it, but the original one is NOT changed
    arr_x[1] = 99
	arr_x[3] = 666
}

=> In comes the slices!

### slices
=============
- Slices in Go are a flexible and efficient way to represent arrays, and they are often used in 
  place of arrays because of their dynamic size and added features.
  Slices in Go are a powerful and flexible data structure that can be used to represent arrays. 
  They provide a more dynamic and efficient way to work with arrays, and they are widely used in Go programs.

- Internally, slice and an array are connected with each other, a slice is a reference to an underlying array.
  

- Go provides several built-in functions that allow you to modify slices, such as append, copy, and delete.

## components of slice
======================
- A slice contains three components:

1. pointer : The pointer is used to points to the first element of the array that is accessible through the slice. 
   Here, it is not necessary that the pointed element is the first element of the array.
   The element x of the array that the slice accesses first, then that will be the pointer of the slice.

2. length: The length is the total number of elements present in the array.

3. Capacity: The capacity represents the maximum size upto which it can expand.
   - Maximum capacity will the size of the underlying array

## Passing Slices to Functions
==============================
=> Slices,unlike the fixed arrays, in contrast, are passed by value but contain a pointer to the original data
Therefore, changes made to a slice inside a function reflect in the original slice.

=> Slices only pass a reference to the underlying array, avoiding duplication of the actual data. 
This makes slices more efficient for large datasets, which is why they’re generally preferred in Go.

=> Only the slice header (pointer, length, and capacity) is passed, not the underlying array data.

func modifySlice(s []int) {
    s[0] = 99
}

func main() {
    mySlice := []int{1, 2, 3, 4, 5}
    modifySlice(mySlice)
    fmt.Println(mySlice) // Output: [99 2 3 4 5], original slice is modified
}

- Efficiency: Passing slices is more efficient than arrays for larger data, as no duplication of the actual data occurs.


*/


package main

import (
   "fmt"
)


func modifySlice(s_x []int){
    // Function takes in a slice
	/*
	Unlike arrays,slices, in contrast, are passed by value but contain a pointer to the original data
    Therefore, changes made to a slice inside a function reflect in the original slice.

	Slices only pass a reference to the underlying array, avoiding duplication of the actual data.

    This makes slices more efficient for large datasets, which is why they’re generally preferred in Go.
	*/
	s_x[1] = 999
	s_x[7] = 666

}



func main() {
    init_s := []int{1, 2, 3, 4, 5, 6 , 7 ,8 ,9}

	fmt.Println("\n###==> Pass slice to a function.")
    modifySlice(init_s)
    fmt.Println(init_s) // Output: [1 999 3 4 5 6 7 666 9].. Slice is changed.
    
}
