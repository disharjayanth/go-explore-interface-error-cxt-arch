package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var errFileDoesNotExists = os.ErrNotExist

func fileOperation(dest, src string) error {
	file, err := os.Open(src)
	defer file.Close()
	if err != nil {
		err = fmt.Errorf("Error while opening file sample1.txt: %w", err)
		return err
	}

	sb, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("Error while reading file sample.txt and converting to slice of byte: %w", err)
		return err
	}

	writeFile, err := os.Create(dest)
	defer writeFile.Close()
	if err != nil {
		err = fmt.Errorf("Error while creating new file sample2.txt: %w", err)
		return err
	}

	err = ioutil.WriteFile("sample2.txt", sb, os.FileMode(os.O_RDWR))
	if err != nil {
		err = fmt.Errorf("Error while writing from sample1.txt to sample2.txt")
		return err
	}

	return nil
}

func main() {
	src := "sample1.txt"
	dest := "sample2.txt"

	err := fileOperation(dest, src)

	if errors.Is(err, errFileDoesNotExists) {
		// Since user's error use fmt
		fmt.Println("You need to type filename that exist (sample1.txt)")
	} else if err != nil {
		log.Println("Error:", err)
		return
	}
}
