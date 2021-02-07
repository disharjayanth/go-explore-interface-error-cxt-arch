package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type FileInfo struct {
	name              string
	FileDoesNotExists error
	FilePathError     error
}

func fileOperation(dest, src string) error {
	file, err := os.Open(src)
	defer file.Close()
	if err != nil {
		err = fmt.Errorf("Error while opening file sample1.txt: %w", err)
		return err
	}

	writeFile, err := os.Create(dest)
	defer writeFile.Close()
	if err != nil {
		err = fmt.Errorf("Error while creating new file sample2.txt: %w", err)
		return err
	}

	_, err = io.Copy(writeFile, file)
	if err != nil {
		return fmt.Errorf("Could not copy file from dest to src: %w", err)
	}

	return nil
}

func main() {
	src := "sample1.txt"
	dest := "sample2.txt"

	var FilePathError *os.PathError
	err := fileOperation(dest, src)

	if errors.Is(err, os.ErrNotExist) && errors.As(err, &FilePathError) {
		// Since user's error(may have entered wrong filename) use fmt
		fmt.Println("Error file you gave does not exist:", FilePathError.Op, "at path:", FilePathError.Path, "operation:", FilePathError.Err)
	} else if errors.As(err, &FilePathError) {
		fmt.Println("Error in copy operation:", FilePathError.Err, FilePathError.Path, FilePathError.Op)
	} else if err != nil {
		log.Println("Error:", err)
		return
	}
}
