package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}

	n, err := io.WriteString(file, "Hello World! This is from main func....")
	if err != nil {
		panic(err)
	}
	fmt.Println("No. of bytes written to file:", n)

	n, err = io.WriteString(file, "More text from main func....")
	if err != nil {
		panic(err)
	}
	fmt.Println("No. of bytes written from main func:", n)

	// Closing the file, so that it can release any resourse
	err = file.Close()
	if err != nil {
		panic(err)
	}
}
