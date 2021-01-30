package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sample.txt")
	// 1) panic(err) => Not good. Since it does not describe what error it is.

	// 2) err = fmt.Errorf("File could not be opended: %w", err)
	// log.Println(err)

	// 3) comparing errors, compare the error you got with error variable in pkg
	// Ex: errors.Is(got_err, typeofError*usually in package variables*)
	// => errors.Is(err, os.ErrPermission) returns bool

	// Old way
	// if err == os.ErrPermission {
	// 	err = fmt.Errorf("You do not have permission to access this file: %w", err)
	// 	log.Println(err)
	// } else if err == os.ErrNotExist {
	// 	err = fmt.Errorf("File does not exist: %w", err)
	// 	log.Println(err)
	// } else if err != nil {
	// 	err = fmt.Errorf("File could not be opended: %w", err)
	// 	log.Panicln(err)
	// }

	// Prefferd way
	if errors.Is(err, os.ErrPermission) {
		err = fmt.Errorf("You do not have permission to open given file: %w", err)
		log.Println(err)
		return
	} else if errors.Is(err, os.ErrNotExist) {
		err = fmt.Errorf("File does not exist: %w", err)
		log.Println(err)
		return
	} else if err != nil {
		err = fmt.Errorf("Could not open file: %w", err)
		log.Println(err)
		return
	}
	defer file.Close()

	fileSliceOfByte, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Could not read from file and make it slice of byte: ", err)
		return
	}

	fmt.Println("Contents of file are: ", string(fileSliceOfByte))
}
