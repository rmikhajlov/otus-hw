package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("offset exceeds file size", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/test_result.txt", 7000, 0)
		require.True(t, errors.Is(err, ErrOffsetExceedsFileSize), "actual err - %v", err)
	})

	t.Run("error while opening file", func(t *testing.T) {
		err := Copy("testdata/not_exist.txt", "testdata/test_result.txt", 0, 0)
		require.Error(t, err, "expected error, got nil")
	})

	t.Run("error while creating file", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/not_exist/test_result.txt", 0, 0)
		require.Error(t, err, "expected error, got nil")
	})

	t.Run("negative offset", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/test_result.txt", -10, 0)
		require.Error(t, err, "expected error, got nil")
	})

	t.Run("negative limit", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/test_result.txt", 0, -10)
		require.Error(t, err, "expected error, got nil")
	})

	t.Run("error while copying the file", func(t *testing.T) {
		err := Copy("testdata", "testdata/test_result.txt", 0, 0)
		require.Error(t, err, "expected error, got nil")
	})

	t.Run("unsupported file", func(t *testing.T) {
		err := Copy("~/dev/urandom", "testdata/test_result.txt", 0, 0)
		require.True(t, errors.Is(err, ErrUnsupportedFile), "actual err - %v", err)
	})
}
