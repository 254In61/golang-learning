# Reading files in Go
The os package provides a ReadFile function that makes reading files straightforward in Go. 

For example, I have a data.txt file in my project folder, and I can read and print it out with the following code:

package main

import (
    "fmt"
    "os"
)

func main() {
    filepath := "data.txt"
    data, err := os.ReadFile(filepath)
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println(string(data))
}