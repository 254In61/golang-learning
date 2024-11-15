/*
## overview
===========
- A goroutine is a lightweight thread of execution in the Go programming language. 
It is a key feature of Go that enables concurrency, allowing a program to perform multiple tasks simultaneously.

## feautures
- The Go runtime manages goroutines, including scheduling and execution, using a mechanism called the Go scheduler.
  Unlike OS threads, goroutines are multiplexed onto a smaller number of OS threads, making them more efficient.

- Goroutines allow for concurrent execution of functions, which means multiple tasks can run independently at the same time.

- When you create goroutine(done by simply prefixing a function call with 'go'), the order will be random.

- main() function has to be kept alive until the goroutine finish. 
  If main() function exits, the goroutine stops.

- Get into Channels and other stuff along the way with projects.
*/

package main

import (
  "fmt"
  "time"
)

func displayFuncA(str_x string){
    fmt.Println("displayFuncA() :",str_x)
}


func displayFuncB(str_x string){
    fmt.Println("displayFuncB() :",str_x)
}

func main(){
    // start a goroutine

    go displayFuncA("Func A Initial String...")
    go displayFuncB("Func B Initial String...")

    names_slc:= []string{"allan", "nancy", "stevo", "albert","allan", "nancy", "stevo", "albert", "nancy", "stevo", "albert","allan", "nancy", "stevo", "albert" }
    
    for x:=0; x < len(names_slc); x++{ 
        go displayFuncA(names_slc[x])
        go displayFuncB(names_slc[x])

    } 


    // keep the main function alive to let the goroutine finish
    time.Sleep(time.Second)
    fmt.Println("Main function ends")
}
