/*

### Overview
==============
- Arrays in Golang or Go programming language is much similar to other programming languages.

- In the program, sometimes we need to store a collection of data of the same type, like a list of student marks.
  Such type of collection is stored in a program using an Array. 

- An array is a FIXED-LENGTH sequence that is used to store homogeneous elements in the memory. 
  Due to their fixed length array are not much popular like Slice in Go language. 
  In an array, you are allowed to store zero or more than zero elements in it. 
  The elements of the array are indexed by using the [] index operator with their zero-based position, 
  which means the index of the first element is array[0] and the index of the last element is array[len(array)-1]. 

### Creating and accessing an Array
====================================
In Go language, arrays are created in two different ways:

1) Using var keyword: 
   var array_name[length]Type
 
2) Shorthand declaration:
   array_name:= [length]Type{item1, item2, item3,...itemN}

### Characteristics of arrays
=================================
- Arrays are mutable.You can change the values of different index position.
- Array type is one-dimensional
- Length of the array is fixed and can't be changed.
- You can store duplicate elements in an array.

### Multi-Dimensional Array
===========================
- Arrays are usully 1-D but you are allowed to create a multi-dimensional array. 
- Multi-Dimensional arrays are the arrays of arrays of the same type. 

Syntax
Array_name[Length1][Length2]..[LengthN]Type

You can create the multi-dimensional array using Var keyword or using shorthand declaration.

## Empty array and default value
================================
- In an array, if an array does not initialize explicitly, then the default value of this array is 0. 

## Length of an array
=====================
- In an array, you can find the length of the array using len() method.
- if ellipsis ‘‘…’’ become visible at the place of length, 
  then the length of the array is determined by the initialized elements.
  
  myarray:= [...]string{"GFG", "gfg", "geeks", "GeeksforGeeks", "GEEK"} 

## value type NOT reference type
================================
- In Go language, an array is of value type not of reference type. 
So when the array is assigned to a new variable, then the changes made in the new variable 
do not affect the original array.

- An array is not something like a pointer to a memory.

## Compare 2 arrays using == operator
- In an array, if the element type of the array is comparable, then the array type is also comparable. 

// Arrays     
arr1:= [3]int{9,7,6} 
arr2:= [...]int{9,7,6} 
arr3:= [3]int{9,5,3} 
  
// Comparing arrays using == operator 
fmt.Println(arr1==arr2)  # outputs true
fmt.Println(arr2==arr3)  # outputs false
fmt.Println(arr1==arr3)  # outputs false


*/
package main

import "fmt"

func main(){
	// Shorthand declaration of array 
    names_arr:= [4]string{"allan", "nancy", "stevo", "albert"} 
	
	// Accessing the elements of  
    // the array Using for loop 
    fmt.Println("\n==> Elements of the name_array:") 
	fmt.Println("Length of the array is : ", len(names_arr))

	for i:=0; i<3; i++{
		fmt.Println(names_arr[i])
	}

	/*
   Output
   $ go run main.go 
   Elements of the array:
   allan
   nancy
   stevo

	*/

	// Multi-dimensional array
	// 2-dimensional array
	language_arr := [3][3]string{{"C#", "C", "Python"}, {"Java", "Scala", "Perl"}, {"C++", "Go", "HTML"}}
	
	// Accessing the values of the 
    // array Using for loop
	fmt.Println("\n==> Elements of language_arr")
	fmt.Println("Length of the array is : ", len(language_arr))

	for x := 0; x < 3; x++ { 
        for y := 0; y < 3; y++ { 
  
            fmt.Println(language_arr[x][y]) 
        } 
    } 

	// empty array
	fmt.Println("\n==> Elements of empty_arr") 
	var empty_arr[2] int 
	fmt.Println("Length of the array is : ", len(empty_arr))
	fmt.Println("Elements of the Array: ", empty_arr) 

	// iterate over the range of the elements of the array.
    fmt.Println("\n==> iterate over the range of the elements of the array.") 
	myarray:= [...]int{29, 79, 49, 39, 20, 49, 48, 49} 

    // Iterate array using for loop 
    for x:=0; x < len(myarray); x++{ 
        fmt.Printf("%d\n", myarray[x]) 
    } 

}