# origins
https://www.geeksforgeeks.org/golang-tutorial-learn-go-programming-language/
https://www.tutorialspoint.com/go/index.htm

- Initially developed at Google in the year 2007 by Robert Griesemer, Rob Pike, and Ken Thompson. But they launched it in 2009 as an open-source programming language.
- Is a statically typed and procedural programming language having syntax similar to C language. 
- It provides:
  -  a rich standard library
  -  garbage collection
  -  type safety
  -  dynamic-typing capability
  -  support for the environment to adopt patterns like dynamic languages. 

- The latest version of the Golang is 1.20.3 released on 4th April 2023.

- Programs are constructed using packages, for efficient management of dependencies. 

- Go programming implementations use a traditional compile and link model to generate executable binaries. 


# Why to Learn Golang?
The main purpose of designing Golang was to eliminate the problems of existing languages. 
So let us see the problems that we are facing with Python, Java, C/C++ programming languages:

- Python: It is easy to use but slow in comparison to Golang.
- Java: It has a very complex type system.
- C/C++: It has slow compilation time as well as a complex type system.

Also, all these languages were designed when multi-threading applications were rare, so not very effective to highly scalable, concurrent and parallel applications.
Threading consumes 1MB whereas Goroutine consumes 2KB of memory, hence at the same time, we can have millions of goroutine triggered.

# Features of Go Programming
- Support for environment adopting patterns similar to dynamic languages. 
  For example, type inference (x := 0 is valid declaration of a variable x of type int)
- Compilation time is fast.
- Inbuilt concurrency support: lightweight processes (via go routines), channels, select statement.
- Go programs are simple, concise, and safe.
- Support for Interfaces and Type embedding.
- Production of statically linked native binaries without external dependencies.

# Features Excluded Intentionally
To keep the language simple and concise, the following features commonly available in other similar languages are omitted in Go:
- Support for type inheritance
- Support for method or operator overloading
- Support for circular dependencies among packages
- Support for pointer arithmetic
- Support for assertions
- Support for generic programming

# Go Compiler
- The source code written in source file is the human readable source for your program.
  It needs to be compiled and turned into machine language so that your CPU can actually execute the program as per the instructions given. 
  
- The Go programming language compiler compiles the source code into its final executable program.

- Go distribution comes as a binary installable for FreeBSD (release 8 and above), Linux, Mac OS X (Snow Leopard and above), 
  and Windows operating systems with 32-bit (386) and 64-bit (amd64) x86 processor architectures.

# installation

$ sudo apt  install -y golang-go

$ go version
go version go1.18.1 linux/amd64

# tokens
- A Go program consists of various tokens. 
- A token is either a keyword, an identifier, a constant, a string literal, or a symbol. 

- For example, the following Go statement consists of six tokens:

fmt.Println("Hello, World!")

The individual tokens are −

fmt
.
Println
(
   "Hello, World!"
)

# line separator
- In a Go program, the line separator key is a statement terminator. That is, individual statements don't need a special separator like “;” in C.
- The Go compiler internally places “;” as the statement terminator to indicate the end of one logical entity.

# identifiers
- A Go identifier is a name used to identify a variable, function, or any other user-defined item. 
- An identifier starts with a letter A to Z or a to z or an underscore _ followed by zero or more letters, underscores, and digits (0 to 9).

identifier = letter { letter | unicode_digit }.

- Go does not allow punctuation characters such as @, $, and % within identifiers. 

- Go is a case-sensitive programming language. Thus, Manpower and manpower are two different identifiers in Go. 

- Here are some examples of acceptable identifiers −

mahesh      kumar   abc   move_name   a_123
myname50   _temp    j      a23b9      retVal

# keywords
- The following list shows the reserved words in Go. 
These reserved words may not be used as constant or variable or any other identifier names.

break	default	func	interface	select
case	defer	Go	map	Struct
chan	else	Goto	package	Switch
const	fallthrough	if	range	Type
continue	for	import	return	Var

# whitespace in Go

Whitespace is the term used in Go to describe blanks, tabs, newline characters, and comments. 
A line containing only whitespace, possibly with a comment, is known as a blank line, and a Go compiler totally ignores it.

Whitespaces separate one part of a statement from another and enables the compiler to identify where one element in a statement, such as int, ends and the next element begins. 
Therefore, in the following statement −

var age int;

There must be at least one whitespace character (usually a space) between int and age for the compiler to be able to distinguish them. 
On the other hand, in the following statement −

fruit = apples + oranges;   // get the total fruit
No whitespace characters are necessary between fruit and =, or between = and apples, although you are free to include some if you wish for readability purpose.

# data types
- Data types refer to an extensive system used for declaring variables or functions of different types. 
  The type of a variable determines how much space it occupies in storage and how the bit pattern stored is interpreted.

- Data types in Go can be classified as follows :

1. Boolean types
   - consists of the two predefined constants: (a) true (b) false

2. Numeric types
   - Arithmetic types and represent :
     - integer types : 
       The predefined architecture-independent integer types are :
       a) uint8  - Unsigned 8-bit integers (0 to 255)
       b) int8   - Signed 8-bit integers (-128 to 127)
       c) uint16 - Unsigned 16-bit integers (0 to 65535)
       d) int16  - Signed 16-bit integers (-32768 to 32767)
       e) uint32 - Unsigned 32-bit integers (0 to 4294967295)
       f) int32  - Signed 32-bit integers (-2147483648 to 2147483647)
       g) uint64 - Unsigned 64-bit integers (0 to 18446744073709551615)
       h) int64  - Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

     - floating types : 
       a) float32    - IEEE-754 32-bit floating-point numbers
       b) float64    - IEEE-754 64-bit floating-point numbers
       c) complex64  - Complex numbers with float32 real and imaginary parts
       d) complex128 - Complex numbers with float64 real and imaginary parts

3. String types
   - A string type represents the set of string values. 
   - Its value is a sequence of bytes. 
   - Strings are immutable types that is once created, it is not possible to change the contents of a string.

4. Derived types
   - They include : 
     (a) Pointer types
     (b) Array types
     (c) Structure types
     (d) Union types 
     (e) Function types 
     (f) Slice types 
     (g) Interface types 
     (h) Map types 
     (i) Channel Types

# variables
## overview
- A variable is nothing but a name given to a storage area that the programs can manipulate. 

- Each variable in Go has a specific type, which determines the size and layout of the variable's memory, 
  the range of values that can be stored within that memory, and the set of operations that can be applied to the variable.

- The name of a variable can be composed of letters, digits, and the underscore character. 
  It must begin with either a letter or an underscore. Upper and lowercase letters are distinct because Go is case-sensitive.

- Variables could be in-built types like :
  - byte
  - int
  - float32

or user defined like :
  - Enumeration
  - Pointer
  - Array
  - Structure
  - Union

## statically typed language
- This means that variable types are known at compile-time rather than at runtime, which helps catch type-related errors early in the development process. 
  In Go, once you declare a variable's type, it cannot change, enforcing strict type checking.

  For example :

  var name string = "John"  // 'name' is declared as a string
  name = 123                // This would result in a compile-time error


## variable definition
- tells the compiler where and how much storage to create for the variable. 
  A variable definition specifies a data type and contains a list of one or more variables of that type as follows:

  var <variable_list> <optional_data_type>;

- Some valid declarations are shown here −

var  i, j, k int;  // The statement “var i, j, k;” declares and defines the variables i, j and k; . This instructs the compiler to create variables named i, j, and k of type int.
var  c, ch byte;
var  f, salary float32;
d =  42;  // Variables can be initialized (assigned an initial value) in their declaration. The type of variable is automatically judged by the compiler based on the value passed to it. 
d = 3, f = 5;    // declaration of d and f. Here d and f are int 

- For definition without an initializer: variables with static storage duration are implicitly initialized with nil (all bytes have the value 0); the initial value of all other variables is zero value of their data type.

## static type declaration
- This means that variable types are known at compile-time rather than at runtime, which helps catch type-related errors early in the development process. 

## Dynamic type declaration / Type inference
- A dynamic type variable declaration requires the compiler to interpret the type of the variable based on the value passed to it. 
  The compiler does not require a variable to have type statically as a necessary requirement.

-  When you declare a variable with :=, Go automatically infers the type based on the assigned value:
   
   name := "John"  // Go infers 'name' as a string

- Despite this convenience, the type remains static once assigned, maintaining Go’s static typing rules. 
  This combination of static typing with type inference is one reason Go is both efficient and easy to work with.

# lvalues vs rvalues
There are two kinds of expressions in Go −

1. lvalue − Expressions that refer to a memory location is called "lvalue" expression. 
   An lvalue may appear as either the left-hand or right-hand side of an assignment.

   Variables are lvalues and so may appear on the left-hand side of an assignment.

   me = me

2. rvalue − The term rvalue refers to a data value that is stored at some address in memory. 
  An rvalue is an expression that cannot have a value assigned to it which means an rvalue may appear on the right- but not left-hand side of an assignment.

  Numeric literals are rvalues and so may not be assigned and can not appear on the left-hand side.

  The following statement is valid: 

  x = 20.0

  The following statement is not valid. It would generate compile-time error:

  10 = 20

# constants
- Constants refer to fixed values that the program may not alter during its execution. 
  These fixed values are also called literals.

- Constants can be of any of the basic data types like an integer constant, 
  a floating constant, a character constant, or a string literal. 
  There are also enumeration constants as well.

- Constants are treated just like regular variables except that their values cannot be modified after their definition

# integer literals
- In Go, integer literals are representations of integer values in the code. 
  They allow you to write numbers in different bases, including decimal, binary, octal, and hexadecimal.

- Types of Integer Literals

1. Decimal (Base 10)

The most common form, where the number is represented in base 10.
No prefix is required.
Example: 42, 123, 0

2. Binary (Base 2)

Represented with a 0b or 0B prefix.
Example: 0b1010 (binary for 10 in decimal)

3. Octal (Base 8)

Represented with a 0 prefix.
Example: 075 (octal for 61 in decimal)

4. Hexadecimal (Base 16)

Represented with a 0x or 0X prefix.
Uses digits 0-9 and letters A-F or a-f.
Example: 0x2A (hexadecimal for 42 in decimal)

# escape sequence

When certain characters are preceded by a backslash, they will have a special meaning in Go. 
These are known as Escape Sequence codes which are used to represent newline (\n), tab (\t), backspace, etc.

\\	\ character
\'	' character
\"	" character
\?	? character
\a	Alert or bell
\b	Backspace
\f	Form feed
\n	Newline
\r	Carriage return
\t	Horizontal tab
\v	Vertical tab
\ooo	Octal number of one to three digits
\xhh . . .	Hexadecimal number of one or more digits

example :

package main

import "fmt"

func main() {
   fmt.Printf("Hello\tWorld!")
}

# operators
An operator is a symbol that tells the compiler to perform specific mathematical or logical manipulations. 
Go language is rich in built-in operators and provides the following types of operators −

1. Arithmetic Operators  : +,-, * , / , % ,++ , --
2. Relational Operators  : ==, != , > , <, >= . <=
3. Logical Operators     : && , || , !
4. Bitwise Operators
5. Assignment Operators  
6. Miscellaneous Operators


