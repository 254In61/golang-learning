package main

import "fmt"

func main(){

	// #### Iterate over a slice
	//==============================
    var sl_names = []string{"Allan", "Nancy", "Shujaa", "Shujaa", "Surry"}
    sl_digits := []int{967, 60, 455, 50,34, 493, 30, 447, 89, 993, 467,777, 888, 333, 672}

	// using for loop
	// simplest way to iterate over a slice
    fmt.Println("\n#### ===> iterating over a slice using for loop")
    for e := 0; e < len(sl_names); e++ {
        fmt.Println(sl_names[e])
    }

    // using range
    // Using range in the for loop, you can get the index and the element value as shown in the example:
	fmt.Println("\n#### ===> iterating over a slice using range")
	for index, ele := range sl_digits {
        fmt.Printf("Index = %d | element = %d\n", index+3, ele)
    }

	// Using a blank identifier in for loop: In the range for loop, if you don’t want to get the index 
	// value of the elements then you can use blank space(_) in place of index variable
	fmt.Println("\n#### ===> iterating over a slice using range with blank identifier")
	
	// Iterate slice sing range in for loop without index
    for _, ele := range sl_names {
        fmt.Printf("Element = %s\n", ele)
    }

	// #### zero value slice
	//=======================
	/*
	In Go language, you are allowed to create a nil slice that does not contain any element in it. 
	So the capacity and the length of this slice is 0. 
	Nil slice does not contain an array reference as shown in the below example:
	*/
    fmt.Println("\n#### ===> zero value slice")
	var myslice []string
    fmt.Printf("Length = %d\n", len(myslice))
    fmt.Printf("Capacity = %d\n", cap(myslice))

	// ### modifying a slice means modifying underlying array
	//=======================================================

	/*
    - As we already know that slice is a reference type it can refer an underlying array. 
	- So if we change some elements in the slice, then the changes should also take place in the referenced array. 
	- Or in other words, if you made any changes in the slice, then it will also reflect in the array as shown in the below example:
	*/
    fmt.Println("\n#### ===> modifying a slice means modifying underlying array")
	// Creating a zero value slice
    arr := [9]int{55, 66, 77, 88, 99, 22, 4 , 51 , 89}
    slc := arr[0:6]
    
	// Before modifying
    fmt.Println("Original_Array: ", arr)
    fmt.Println("Original_Slice: ", slc)

    // After modification
    slc[0] = 100
    slc[2] = 222
    slc[5] = 555

	fmt.Println("\nModified_Array: ", arr)
    fmt.Println("Modified_Slice: ", slc)

	// ### Comparison of Slice
	// ============================

	// You can check if slice is nil or not and get a bool output. 
	// You CANNOT compare 2 slices though, you'll get a compiler error.
    
	fmt.Println("\n#### ===> Comparison of Slice")
	// creating slices
    s1 := []int{12, 34, 56}
    var s2 []int
	// Checking if the given slice is nil or not
    fmt.Println(s1 == nil)
    fmt.Println(s2 == nil)
    
	/* 
	- the below instructions will cause a compiler error.
	s3:= []int{23, 45, 66}
    fmt.Println(s1==s3)
    
	$ go run main.go 
    # command-line-arguments
    ./main.go:87:17: invalid operation: s1 == s3 (slice can only be compared to nil)

	Note: If you want to compare two slices, then use range for loop to match each element or you can use DeepEqual function. 
    */
     
	
	// ### Multi-dimensional slice
	// Multi-dimensional slice are just like the multidimensional array, except that slice does not contain the size.
    // Creating multi-dimensional slice

	// A multi-dimensional slice of integers ([][]int). 
	// Each inner slice (e.g., {12, 34}) represents a row in the multi-dimensional structure.
    s1 := [][]int{
		{12, 34},
        {56, 47},
        {29, 40},
        {46, 78},
    }
 
    // Accessing multi-dimensional slice
    fmt.Println("Slice 1 : ", s1)

    fmt.Println("\n#### ===> Multi-dimensional slice")
	s2 := [][]string{
        []string{"Geeks", "for"},
        []string{"Geeks", "GFG"},
        []string{"gfg", "geek"},
    }
    
	// Accessing multi-dimensional slice
    fmt.Println("Slice 2 : ", s2)
	


}