package main

import (
	"errors"
	"fmt"
	"time"
)

// ErrFileNotFound is custom error type
type ErrFileNotFound struct {
	FileName string
	When     time.Time
}

func (e ErrFileNotFound) Error() string {
	return fmt.Sprintf("Given file %s is not found at time %v\n", e.FileName, e.When)
}

// func (e ErrFileNotFound) Is(err error) bool {
// 	_, ok := err.(ErrFileNotFound)
// 	return ok
// }

func main() {
	errFile := ErrFileNotFound{
		FileName: "sample.txt",
		When:     time.Now(),
	}

	is := errors.Is(errFile, ErrFileNotFound{
		FileName: "sample.txt",
		When:     errFile.When,
	})

	fmt.Println(is)

	fmt.Println(errFile)
}
