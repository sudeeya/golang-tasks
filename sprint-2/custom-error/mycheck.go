//go:build !solution

package mycheck

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

const maxStringLength int = 19

var (
	ErrContainsNumber = errors.New("found numbers")
	ErrIsTooLong      = errors.New("line is too long")
	ErrNoTwoSpaces    = errors.New("no two spaces")
)

type ErrorList []error

func (el ErrorList) Error() string {
	errorStrings := make([]string, len(el))
	for i, e := range el {
		errorStrings[i] = e.Error()
	}
	return strings.Join(errorStrings, ";")
}

func MyCheck(input string) error {
	var (
		errs           []error
		containsNumber bool
		spaceCount     int = 0
	)
	for _, r := range input {
		if !containsNumber && unicode.IsNumber(r) {
			containsNumber = true
		} else if unicode.IsSpace(r) {
			spaceCount++
		}
	}
	if containsNumber {
		errs = append(errs, ErrContainsNumber)
	}
	if utf8.RuneCountInString(input) > maxStringLength {
		errs = append(errs, ErrIsTooLong)
	}
	if spaceCount != 2 {
		errs = append(errs, ErrNoTwoSpaces)
	}
	return ErrorList(errs)
}
