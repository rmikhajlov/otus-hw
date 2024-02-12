package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if offset < 0 || limit < 0 {
		return fmt.Errorf("offset and limit should be >= 0")
	}

	fileFrom, err := os.OpenFile(fromPath, os.O_RDONLY, 0)
	if err != nil {
		return ErrUnsupportedFile
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
		return fmt.Errorf("error while seeking file: %w", err)
	}

	fileTo, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("error while creating file: %w", err)
	}
	defer fileTo.Close()

	var pbLimit int64

	if limit+offset > fileFromStat.Size() {
		pbLimit = fileFromStat.Size() - offset
	} else {
		pbLimit = limit
	}

	bar := pb.Full.Start64(pbLimit)

	barReader := bar.NewProxyReader(fileFrom)

	if pbLimit == 0 || pbLimit > fileFromStat.Size() {
		_, err = io.Copy(fileTo, barReader)
		if err != nil {
			os.Remove(toPath)
			return fmt.Errorf("error while copying file: %w", err)
		}
	} else {
		_, err = io.CopyN(fileTo, barReader, pbLimit)
		if err != nil {
			os.Remove(toPath)
			return fmt.Errorf("error while copying file: %w", err)
		}
	}

	bar.Finish()

	return err
}
