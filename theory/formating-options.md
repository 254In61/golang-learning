# Overview
In Golang, you can format strings and other data types using the fmt package, which provides a variety of verbs and options to control the output.

# Basic Formatting with Verbs
%v: The default format, which varies depending on the type.
%+v: Similar to %v but includes field names in structs.
%#v: A Go-syntax representation of the value (useful for debugging).
%T: Prints the type of the value

fmt.Printf("%v\n", 42)             // 42
fmt.Printf("%+v\n", struct{ Name string }{"John"}) // {Name:John}
fmt.Printf("%#v\n", struct{ Name string }{"John"}) // struct { Name string }{Name:"John"}
fmt.Printf("%T\n", 42)             // int


# Integer Formatting
%d: Decimal.
%b: Binary.
%o: Octal.
%x / %X: Hexadecimal (lowercase / uppercase).
%c: Character corresponding to the Unicode code point.

fmt.Printf("%d\n", 42)      // 42
fmt.Printf("%b\n", 42)      // 101010
fmt.Printf("%o\n", 42)      // 52
fmt.Printf("%x\n", 42)      // 2a
fmt.Printf("%X\n", 42)      // 2A
fmt.Printf("%c\n", 65)      // A

# Floating-Point Formatting
%f: Decimal point, no exponent (e.g., 123.456).
%e / %E: Scientific notation (e.g., 1.23456e+02).
%g / %G: Automatically chooses %f or %e based on the value.

fmt.Printf("%f\n", 123.456)    // 123.456000
fmt.Printf("%.2f\n", 123.456)  // 123.46 (two decimal places)
fmt.Printf("%e\n", 123.456)    // 1.234560e+02
fmt.Printf("%g\n", 123.456)    // 123.456

# String and Character Formatting
%s: String.
%q: Double-quoted string (includes escape characters).
%x: Hex dump (2 characters per byte).
% X: Hex dump with spaces between bytes

fmt.Printf("%s\n", "Golang")         // Golang
fmt.Printf("%q\n", "Golang")         // "Golang"
fmt.Printf("%x\n", "Golang")         // 476f6c616e67
fmt.Printf("% X\n", "Golang")        // 47 6f 6c 61 6e 67

# Boolean Formatting
%t: Prints true or false.

fmt.Printf("%t\n", true)   // true
fmt.Printf("%t\n", false)  // false

# Pointer Formatting
%p: Prints the memory address of a pointer.

var a int = 10
fmt.Printf("%p\n", &a)  // 0xc0000b2000 (address will vary)

# Width and Precision Control
Width: Specify minimum width with %[width]v. For example, %5d will print a number right-aligned within a 5-character space.
Precision: Control precision with %.precision, like %.2f for two decimal places.
Width and Precision Together: Combine width and precision, e.g., %6.2f.

fmt.Printf("%6d\n", 42)        //     42
fmt.Printf("%6.2f\n", 3.14159) //   3.14

# Using fmt.Sprintf for Formatted Strings

If you need to store formatted output in a string, use fmt.Sprintf instead of fmt.Printf.

formattedString := fmt.Sprintf("Hello, %s!", "Golang")
fmt.Println(formattedString)  // Hello, Golang!

# Summary of common verbs
%v	Default format	             fmt.Sprintf("%v", 42)
%+v	Struct with field names	     fmt.Sprintf("%+v", struct{Name string}{Name: "Alice"})
%#v	Go-syntax representation	 fmt.Sprintf("%#v", struct{Name string}{Name: "Alice"})
%T	Type of the value	         fmt.Sprintf("%T", 42)
%d	Integer (decimal)	         fmt.Sprintf("%d", 42)
%f	Float	                     fmt.Sprintf("%f", 3.14)
%s	String	                     fmt.Sprintf("%s", "Golang")
%q	Quoted string	             fmt.Sprintf("%q", "Golang")
%p	Pointer	                     fmt.Sprintf("%p", &a)
%t	Boolean	                     fmt.Sprintf("%t", true)

