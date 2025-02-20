package users

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(input string) (string, error) {
	buf := bytes.NewBufferString(input)
	if buf.Len() > 72 {
		return "", fmt.Errorf("input string is longer than 72 bytes")
	}

	hash, err := bcrypt.GenerateFromPassword(buf.Bytes(), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
