package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("file01.txt", os.O_RDWR, 0755)
	if err != nil {
		fmt.Println("Cannot open file01.txt", err)
	}

	newFile, err := os.Create("file02.txt")
	if err != nil {
		fmt.Println("Cannot create new file file02.txt", err)
	}

	written, err := io.Copy(newFile, file)
	if err != nil {
		fmt.Println("Cannot write to newly created file from another file")
	}

	fmt.Println("No. of bytes written: ", written)
}
