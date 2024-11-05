package main

import "fmt"

func main(){
	var x int
	x = 23
    fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)

	var y float64 = 27.34
	fmt.Println(y)
	fmt.Printf("y is of type %T\n", y)
    
	// initialize value with := operator. This is inference
	// Variable has been declared without any type
	r := 47
	fmt.Println(r)
	fmt.Printf("r is of type %T\n", r)


    u := "Zimbabwe"
	fmt.Println(u)
	fmt.Printf("u is of type %T\n", u)

	// Mixed variables
	// Variables of different types can be declared in one go using type inference.

   var a, b, c = 3, 4, "foo"  
	
   fmt.Println(a)
   fmt.Printf("a is of type %T\n", a)

   fmt.Println(b)
   fmt.Printf("b is of type %T\n", b)

   fmt.Println(c)
   fmt.Printf("c is of type %T\n", c)
}