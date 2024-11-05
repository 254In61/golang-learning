/*
- 1st line of the program package defines the package name in which the program should lie.
- Since Go programs run in packages, this is a mandatory statement.
- The main package is the starting point to run the program. 
- Each package has a path and name associated with it.
*/
package main


/*
A preprocessor command which tells the Go compiler to include the files lying in the package fmt.
*/
import "fmt"

func main(){
    /*
    - fmt package has exported Println method which is used to display the message on the screen.
	- Notice the capital P of Println method. 
	  In Go language, a name is exported if it starts with capital letter. 
	  Exported means the function or variable/constant is accessible to the importer of the respective package.
	*/
	fmt.Println("Hello Hello!\nBack to this again!")
}