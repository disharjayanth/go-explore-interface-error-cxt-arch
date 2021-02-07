package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func fileOperation(dest, src string) error {
	file, err := os.Open("sample1.txt")
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

	writeFile, err := os.Create("sample2.txt")
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
	if err != nil {
		log.Println("Error:", err)
		return
	}
}
