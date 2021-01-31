package main

import (
	"fmt"
	"io"
	"os"
)

// WriteFile has file and error type as field
type WriteFile struct {
	f   *os.File
	err error
}

// NewWriteFile returns pointer to writeFile
func NewWriteFile(name string) *WriteFile {
	f, err := os.Create(name)
	return &WriteFile{
		f:   f,
		err: fmt.Errorf("Error while creating new file: %w", err),
	}
}

// WriteString takes in content and writes it to file
func (wf *WriteFile) WriteString(content string) {
	if wf.err != nil {
		return
	}
	_, err := io.WriteString(wf.f, content)
	wf.err = fmt.Errorf("Error while writing string to file: %w", err)
}

// Close closes the file
func (wf *WriteFile) Close() {
	if wf.err != nil {
		return
	}
	err := wf.f.Close()
	if err != nil {
		wf.err = fmt.Errorf("Error while closing file: %w", err)
	}
}

// Error returns error if there is any
func (wf *WriteFile) Error() error {
	return wf.err
}

func main() {
	wf := NewWriteFile("sample.txt")
	wf.WriteString("Hello !")
	wf.WriteString("More Text!")
	wf.Close()

	err := wf.Error()
	if err != nil {
		err = fmt.Errorf("Error in file opertion: %w", err)
		fmt.Println(err)
	}
}
