/*
### overview
============
=> A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly
different types into a single type.
=> Any real-world entity which has some set of properties/fields can be represented as a struct. 
=> This concept is generally compared with the classes in object-oriented programming.
   It can be termed as a lightweight class that does not support inheritance but supports composition. 

=> For Example, an address has a name, street, city, state, Pincode. 
   It makes sense to group these three properties into a single structure address as shown below.

    type Address struct {
      name string 
      street string
      city string
      state string
      Pincode int
}

=> The 'type' keyword introduces a new type. 
   It is followed by the name of the type (Address) and the keyword struct to illustrate that we’re defining 
   a struct. The struct contains a list of various fields inside the curly braces. 
   Each field has a name and a type. 

   We can also make them compact by combining the various fields of the same type as shown in the below example:

type Address struct {
    name, street, city, state string
    Pincode int
}

=> Fields can be accessed using the dot . operator.

=> Methods can be defined on structures using the receiver syntax.

### syntax
===============


var a Address

- The above code creates a variable of a type Address which is by default set to zero.
  For a struct, zero means all the fields are set to their corresponding zero value.
  So the fields name, street, city, state are set to “”, and Pincode is set to 0. 

- You can also initialize a variable of a struct type using a struct literal as shown below:
  var a = Address{"Akshay", "PremNagar", "Dehradun", "Uttarakhand", 252636}

=> 2 aspects :

1. Always pass the field values in the same order in which they are declared in the struct. 
   Also, you can’t initialize only a subset of fields with the above syntax.

2. Go also supports the name: value syntax for initializing a struct 
(the order of fields is irrelevant when using this syntax). 
And this allows you to initialize only a subset of fields. 
All the uninitialized fields are set to their corresponding zero value. 

Example:
var a = Address{Name:”Akshay”, street:”PremNagar”, state:”Uttarakhand”, Pincode:252636} //city:””

### advantages of structures in Go
==================================
1. Encapsulation : Structures allow you to encapsulate related data together, making it easier to manage and modify the data.
2. Code organization: Structures help to organize code in a logical way, which makes it easier to read and maintain.
3. Flexibility: Structures allow you to define custom types with their own behavior, making it easier to work with complex data.
4. Type safety: Structures provide type safety by allowing you to define the type of each field, which helps to prevent 
   errors caused by assigning the wrong type of value.
5. Efficiency: Structures in Go are very efficient, both in terms of memory usage and performance.

### disadvantages of using structures in Go:
============================================
1. Complexity: Structures can make code more complex, especially if the structures have a large number of fields or methods.
2. Boilerplate code: When defining large structures with many fields, it can be time-consuming to write out all of the field names and types.
3. Inheritance: Go does not support inheritance, which can make it more difficult to work with large hierarchies of related data.
4. Immutability: Go structures are mutable by default, which can make it more difficult to enforce immutability in your code.

=> Overall, the advantages of using structures in Go typically outweigh the disadvantages, 
   as they provide a powerful tool for managing and working with complex data. 
   However, as with any programming technique, it’s important to use structures judiciously and be aware of their limitations.

*/

package main

import (
   "fmt"
)

// Defining a struct type
type Person struct {
    Name    string
    Age     int
    Gender  string
    Weight  float32
}


func main() {
    // Declaring a variable of a `struct` type
    // All the struct fields are initialized 
    // with their zero value
    fmt.Println("\n####==> declaring a struct")
    var p1 Person
    fmt.Println("Person1 :",p1)

    /*
    Output
    =======
    $ go run main.go 
      { 0  0}    // strings zero output i.e "" is here because it's an empty space!

    */


    // Declaring and initializing a
    // struct using a struct literal

    p2 := Person{"Allan",42, "Male", 98.5}
    fmt.Println("Person3 :",p2)

    /*
    Output
    =======
    Person2 : {Allan 42 Male 98.5}
    */


    // Uninitialized fields are set to their corresponding zero-value
    p3 := Person{Name: "Maseghe", Weight: 98.51}
    fmt.Println("Person3 :",p3)

    /*
    Output
    =======
    Person3 : {Maseghe 0  98.51}
    */
    
    // Accessing struct fields using the dot operator
    fmt.Println("\n####==> accessing struct fields")
    fmt.Println("Person3 Name   : ", p3.Name)
    fmt.Println("Person3 Age    : ", p3.Age)
    fmt.Println("Person3 Gender : ", p3.Gender)
    fmt.Println("Person3 Weight : ", p3.Weight)

    /*
    output
    =======
    Person3 Name   :  Maseghe
    Person3 Age    :  0
    Person3 Gender :            // Blank since we left the gender blank
    Person3 Weight :  98.51
    */
}
