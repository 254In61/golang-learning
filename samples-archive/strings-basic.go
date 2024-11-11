/*
See more string functions at the bottom of this page : https://www.geeksforgeeks.org/strings-in-golang/?ref=shm
** All I need are basics, then I get into building project!.. Complexities to be learnt as I am flying!**

- strings serve the purpose of representing text

Immutability
============
- Strings are immutable i.e once a string is created, it cannot be modified.
- Any operation that seems to "modify" a string actually creates a new string.

Unicode and Encoding
====================
- Strings are UTF-8 encoded by default, and each string is a read-only slice of bytes. 
- This makes it easy to handle both ASCII and non-ASCII characters efficiently. 
  Due to UTF-8 encoding Golang string can contain a text which is the mixture of 
  any language present in the world, without any confusion and limitation of the page.
- However, accessing a string by index yields a byte, not a character. 
  To iterate over characters (runes) in a string, Go provides the for range loop.

Accessing Characters
====================
- Accessing a string by index (s[i]) returns a byte representing the UTF-8 encoded value. 
 To get a character (rune), you can use a loop or convert the byte as needed.

 s := "Golang"
 fmt.Println(s[0])       // Outputs: 71 (byte value of 'G')
 fmt.Println(string(s[0])) // Outputs: G (converted to string)

String Concatenation
=====================
- Golang uses the + operator for concatenation. 
- Strings are immutable, so concatenation results in a new string. 
- For efficient concatenation, especially in loops, Go provides the strings.Builder type.

Multiline Strings
==================
- Multiline strings can be created using backticks (`), which ignore escape characters, making them useful for raw strings.

String Formatting
=================
- Golang uses fmt.Sprintf for string formatting with format specifiers, similar to printf in C. 
- Go does not support f-strings or interpolation within raw strings.

String Slicing
===============
- Slicing a string returns a substring, which is still immutable. 
- Syntax 
  s[start:end]

*/


package main

import (
   "fmt"
   "unicode/utf8"
)

func display(str_x string){
	fmt.Println("String : ",str_x)

	// Iterate over the string using for range loop
	for index, item := range str_x{
		fmt.Printf("Index: %d | Value: %c\n", index, item)
	}

}

func main(){
   // #### Creating a string
   
   var names,countries string
   _ = countries  // A way to go past the error, ./main.go:68:14: countries declared but not used
   
   fmt.Println("#### ===> Creating a string")
   names = "Allan Maseghe"
   // display(names)
   
   /*
   This instruction will lead to compiler throwing an error.
   Why? Because strings are immutable. 

   names[3] = 'll'
   */

   // #### access the individual byte of the string
   // The string is of a byte so, we can access each byte of the given string.
   // Accessing the bytes of the given string
   for c := 0; c < len(names); c++ {
	 fmt.Printf("\nCharacter = %c Bytes = %x\n", names[c], names[c])
   }
   
   /*
  Character = A Bytes = 41
  Character = l Bytes = 6c
  Character = l Bytes = 6c
  Character = a Bytes = 61
  Character = n Bytes = 6e
  Character =   Bytes = 20
  Character = M Bytes = 4d
  Character = a Bytes = 61
  Character = s Bytes = 73
  Character = e Bytes = 65
  Character = g Bytes = 67
  Character = h Bytes = 68
  Character = e Bytes = 65

   */

   // #### Find length of a string
   /*
   - In Golang string, you can find the length of the string using two functions:

    a) len() : returns the number of bytes of the string.
	b) RuneCountInString().: provided by UTF-8 package, this function returns the total number of rune presents in the string. 
   */
   
   // using len() function
   fmt.Println("Using len() the Length is:", len(names))

   // using RuneCountInString() function
   fmt.Println("Using len() the Length is:", utf8.RuneCountInString(names))
	

}