package filewriter

import (
	"io"
	"os"
)

type WriteFileError struct {
	Op  string
	Err error
}

func (w WriteFileError) Error() string {
	return w.Err.Error()
}

func (w WriteFileError) Unwrap() error {
	return w.Err
}

type WriteFile struct {
	f   *os.File
	err error
}

func NewFileWrite(name string) *WriteFile {
	file, err := os.Create(name)
	if err != nil {
		return &WriteFile{
			f: nil,
			err: WriteFileError{
				Op:  "Error while creating file",
				Err: err,
			},
		}
	}

	return &WriteFile{
		f:   file,
		err: nil,
	}
}

func (wf *WriteFile) Write(content string) {
	if wf.err != nil {
		return
	}

	_, err := io.WriteString(wf.f, content)

	if err != nil {
		wf.err = WriteFileError{
			Op:  "Error while writing to file",
			Err: err,
		}
	}
}

func (wf *WriteFile) Close() {
	if wf.f == nil {
		return
	}

	err := wf.f.Close()
	if err != nil {
		wf.err = WriteFileError{
			Op:  "Error while closing file",
			Err: err,
		}
	}
}

func (wf *WriteFile) Error() error {
	return wf.err
}
