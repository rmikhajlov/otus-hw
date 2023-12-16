package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var b strings.Builder

	if len(str) == 0 {
		return "", nil
	}

	if _, err := strconv.Atoi(string(str[0])); err == nil {
		return "", ErrInvalidString
	}

	for i := 0; i < len(str); i++ {
		if i < len(str)-1 {
			if str[i+1] == '0' {
				continue
			}
		}
		reps, err := strconv.Atoi(string(str[i]))
		if err == nil {
			if _, err = strconv.Atoi(string(str[i-1])); err == nil {
				return "", ErrInvalidString
			}
			if reps != 0 {
				fmt.Fprintf(&b, "%s", strings.Repeat(string(str[i-1]), reps-1))
			}
		} else {
			fmt.Fprintf(&b, "%s", string(str[i]))
		}
	}

	return b.String(), nil
}
