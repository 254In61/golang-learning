/*
### overview
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

## how to create and initialize a slice
===========================================
1) Using a slice lateral
    - The creation of slice literal is just like an array literal, 
     but with one difference you are not allowed to specify the size of the slice in the square braces[]. 

	 var my_slice_1 = []string{"Geeks", "for", "Geeks"}

    - Always remember when you create a slice using a string literal,
	  it first creates an array and after that return a slice reference to it.

2) Using an Array
    - We already know that  slice is the reference of the array so you can create a slice from the given array.
    - For creating a slice from the given array first you need to specify the lower and upper bound, 
     which means slice can take elements from the array starting from the lower bound to the upper bound. 
	 It does not include the elements above from the upper bound.

3) Using already Existing Slice
    - For creating a slice from the given slice first you need to specify the lower and upper bound, 
	  which means slice can take elements from the given slice starting from the lower bound to the upper bound. 
	  It does not include the elements above from the upper bound.

	- Can be in different flavors 
    var my_slice_1 = oRignAl_slice[1:5]
    my_slice_2 := oRignAl_slice[0:]     # Start from the index[0] to last element.. replicates the original
    my_slice_3 := oRignAl_slice[:6]     # starting from the index[0] element to index[x-1] i,e index[6-1=5] element
    my_slice_4 := oRignAl_slice[:]      # everything replicated
    my_slice_5 := my_slice_3[2:4]

4) Using make() function
    - You can also create a slice using the make() function which is provided by the go library. 
	- This function takes three parameters, i.e, type, length, and capacity. 
	- Here, capacity value is optional. 
	  It assigns an underlying array with a size that is equal to the given capacity and returns a 
	  slice which refers to the underlying array. 
	- Generally, make() function is used to create an empty slice. 
	  Here, empty slices are those slices that contain an empty array reference.

	  func make([]T, len, cap) []T

*/
package main

import "fmt"

func display_str_slice(sx []string){
   fmt.Println("Slice : ", sx)
   fmt.Printf("Length of the slice: %d\n", len(sx))
   fmt.Printf("Capacity of the slice: %d\n", cap(sx))
   fmt.Printf("Pointer of the slice: %p\n", &sx) // address of the slice header, not the underlying array.
   fmt.Printf("Pointer to the underlying array: %p\n", &sx[0]) // address of the underlying array that the slice references
   fmt.Printf("\n")
}

func display_int_slice(sx []int){
	fmt.Println("Slice : ", sx)
	fmt.Printf("Length of the slice: %d\n", len(sx))
	fmt.Printf("Capacity of the slice: %d\n", cap(sx))
	fmt.Printf("Pointer of the slice: %p\n", &sx) // address of the slice header, not the underlying array.
	fmt.Printf("Pointer to the underlying array: %p\n", &sx[0]) // address of the underlying array that the slice references
	fmt.Printf("\n")
 }

func main(){

	// #### 1. Create a slice using a slice lateral
    fmt.Println("#### ===> create a slice using a slice lateral")
	var sl_names = []string{"Allan", "Nancy", "Shujaa", "Shujaa", "Surry"}
	display_str_slice(sl_names)

    // #### 2. Create a slice from an array
	// create an array
	africa_arr := [5]string{"Kenya", "Uganda", "Rwanda","Botswana","Tanzania"}
    
	// #### 3. create a slice from the array
	// You declare with Shorthand declaration: :=
	fmt.Println("#### ===> Create slice from an array")
    sl_countries := africa_arr[1:4] // [a:b] => from index[a] to index[b-1]
	display_str_slice(sl_countries)
    
    // #### 4. Creating a slice from existing slice
	// look at the underlying array pointer outputs.. When a new array is formed, that's when pointer changes
	fmt.Println("#### ===> Creating a slice from existing slice")
    original_slice := []int{967, 60, 455, 50,34, 493, 30, 447, 89, 993, 467,777, 888, 333, 672}
    display_int_slice(original_slice )

	sl_1 := original_slice[0:] // underlying array doesn't change from the original..pointer for underlying array is the same
	display_int_slice(sl_1)

    sl_2 := original_slice[3:7] // underlying array changes change from the original..therefore pointer for underlying array changes
	display_int_slice(sl_2)

	sl_3 := original_slice[:] // underlying array doesn't change from the original..pointer for underlying array is the same
	display_int_slice(sl_3)
	
	// #### 5. create a slice using make() function
	fmt.Println("#### ===> Create a slice using make() function")
	// Creating an array of size 7 , slice this array till 4 , return the reference of the slice Using make() function
	var sl_a = make([]int, 4, 7)
	display_int_slice(sl_a)
    
	// Creating an array of size 7 and return the reference of the slice Using make() function
	var sl_b = make([]int, 7)
	display_int_slice(sl_b)
    
    
}