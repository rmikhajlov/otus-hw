package main

import (
	"errors"
	"fmt"
	"github.com/cheggaaa/pb/v3"
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

	var pbLimit int64

	if limit == 0 {
		pbLimit = fileFromStat.Size() - offset
	} else {
		pbLimit = limit - offset
	}

	bar := pb.Full.Start64(pbLimit)

	barReader := bar.NewProxyReader(fileFrom)

	if limit == 0 || limit > fileFromStat.Size() {
		_, err = io.Copy(fileTo, barReader)
		if err != nil {
			return fmt.Errorf("error while copying file: %v", err)
		}
	} else {
		_, err = io.CopyN(fileTo, barReader, limit)
		if err != nil {
			return fmt.Errorf("error while copying file: %v", err)
		}
	}

	bar.Finish()

	return err
}
