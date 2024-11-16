package main

import(
	"fmt"
	"os"
)

func main(){
	file_path := "files/nofile.txt" // No such file so expecting some error.
	data, err := os.ReadFile(file_path)

	if err != nil {
		fmt.Println("File reading error", err)
        return
	}

	fmt.Println(string(data))
}