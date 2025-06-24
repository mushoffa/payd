package valueobject

import (
	// "fmt"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	PREFIX_2a = "$2a$"
	PREFIX_2b = "$2b$"
	PREFIX_2y = "$2y$"
)

var (
	ErrEmptyString = errors.New("password text cannot be empty string")
)

type Password string

func (e Password) String() string {
	return string(e)
}

func (e Password) IsHashed() bool {
	if e.isEmpty() {
		return false
	}

	return isHashed(e.String())
}

func (e Password) Hash() (string, error) {
	if e.isEmpty() {
		return "", ErrEmptyString
	}

	if e.IsHashed() {
		return e.String(), nil
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(e), bcrypt.DefaultCost)
	return string(hashed), err
}

func (e Password) Compare(compareText string) (bool, error) {
	if (e.isEmpty()) || (compareText == "") {
		return false, nil
	}

	var state int = 0b00

	if e.IsHashed() {
		state = state | 0b10
	}

	if isHashed(compareText) {
		state = state | 0b01
	}

	return e.compareState(state, compareText)
}

func (e Password) compareState(state int, compareText string) (bool, error) {
	switch state {
	// Both password and input text are in the same state
	// Either both texts are in plaintext or hashed
	case 0, 3:
		return (e.String() == compareText), nil

	// Password is plaintext, compare text is encypted
	case 1:
		if err := bcrypt.CompareHashAndPassword([]byte(compareText), []byte(e.String())); err != nil {
			return false, err
		}

		return true, nil

	// Password is hashed, compareText is in plaintext
	case 2:
		if err := bcrypt.CompareHashAndPassword([]byte(e.String()), []byte(compareText)); err != nil {
			return false, err
		}

		return true, nil

	default:
		return false, errors.New("invalid state condition")
	}
}

func (e Password) isEmpty() bool {
	return e.String() == ""
}

func isHashed(text string) bool {
	return hasPrefix(text, PREFIX_2a) || hasPrefix(text, PREFIX_2b) || hasPrefix(text, PREFIX_2y)
}

func hasPrefix(text, prefix string) bool {
	return strings.HasPrefix(text, prefix)
}
