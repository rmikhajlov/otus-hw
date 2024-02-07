package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {

	fileFrom, err := os.OpenFile(fromPath, os.O_RDONLY, 0)
	if err != nil {
		return fmt.Errorf("error while opening file: %v", err)
	}
	defer fileFrom.Close()

	fileFromStat, err := fileFrom.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}

	if offset > fileFromStat.Size() {
		return ErrOffsetExceedsFileSize
	}

	_, err = fileFrom.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("error while seeking file: %v", err)
	}

	fileTo, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("error while creating file: %v", err)
	}
	defer fileTo.Close()

	if limit == 0 || limit > fileFromStat.Size() {
		_, err = io.Copy(fileTo, fileFrom)
	} else {
		fmt.Println("COPIED PART")
		_, err = io.CopyN(fileTo, fileFrom, limit)
	}

	return err
}
