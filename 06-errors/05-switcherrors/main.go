package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	// error since sample.txt is not present
	file, err := os.Open("sample.txt")
	defer file.Close()

	var perr *os.PathError

	switch {
	case errors.Is(err, os.ErrPermission) && errors.As(err, &perr):
		err = fmt.Errorf("Case 1:You do not have permission to open file: %w", err)
		log.Println(err)
	case errors.Is(err, os.ErrNotExist) && errors.As(err, &perr):
		err = fmt.Errorf("Case 2:File does not exist: %w", err)
		log.Println(err)
	case errors.As(err, &perr):
		err = fmt.Errorf("Case 3:Here is orginal error: %w and here is path from type PathError after assesrtion: %s", err, perr.Path)
		log.Println(err)
	case err != nil:
		log.Println("Case 4:", err)
	}
}
