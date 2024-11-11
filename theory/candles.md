# imports  
https://www.geeksforgeeks.org/import-in-golang/?ref=oin_asr15

# blanks  - Unused meodules and variables
=> Go compiler raises an error whenever there's imported package or a created variable that is not used anywhere.

=> But many a times you will build a code with an end game in mind..You will declare variables that won't be built immediately
  as you are testing..Or imported packages.

=> You get around that with _

import _ "math" 

or

import ( 
    "fmt"
    _ "math/rand"
) 

variables :

var names,countries string
   
_ = countries  // A way to go past the error, ./main.go:68:14: countries declared but not used
   
fmt.Println("#### ===> Creating a string")
names = "Allan Maseghe"

# functions
func function_name(Parameter-list)(Return_type) {
    // function body...
}

## call function by value
func multiply(a, b int) int {
    a = a * 2 // modifying a inside the function
    return a * b
}

func main() {
    x := 5
    y := 10
    fmt.Printf("Before: x = %d, y = %d\n", x, y)
    result := multiply(x, y)
    fmt.Printf("multiplication: %d\n", result)
    fmt.Printf("After: x = %d, y = %d\n", x, y)
}

