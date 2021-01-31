package main

import (
	"fmt"

	"github.com/disharjayanth/go-explore-interface-error-cxt-arch/06-errors/10-pkgerrors/filewriter"
)

func main() {
	wf := filewriter.NewFileWrite("sample.txt")
	wf.Write("Hi there!")
	wf.Write("Hello....")

	wf.Close()

	err := wf.Error()
	if err != nil {
		err = fmt.Errorf("Error in file operation: %w", err)
		return
	}
}
