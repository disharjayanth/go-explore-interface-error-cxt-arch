package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var errPath *os.PathError

	// sample.txt
	file, err := os.Open("sampl.txt")
	defer file.Close()
	if errors.As(err, &errPath) {
		err = fmt.Errorf("Orginal error err: %w and path from errPath: %s", err, errPath.Path)
		fmt.Println(err)
		fmt.Println("Error of type os.PathError:", errPath.Path, errPath.Op, errPath.Err)
		return
	}
}
