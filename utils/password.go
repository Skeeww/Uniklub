package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(input string) (string, error) {
	buf := []byte(input)
	if buf == nil {
		return "", fmt.Errorf("input string to buffer is nil")
	}
	if len(buf) > 72 {
		return "", fmt.Errorf("input string is longer than 72 bytes")
	}

	hash, err := bcrypt.GenerateFromPassword(buf, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyPassword(input string, hash string) error {
	inputBuf := []byte(input)
	if inputBuf == nil {
		return fmt.Errorf("input string to buffer is nil")
	}

	hashBuf := []byte(hash)
	if hashBuf == nil {
		return fmt.Errorf("hash string to buffer is nil")
	}

	if err := bcrypt.CompareHashAndPassword(hashBuf, inputBuf); err != nil {
		return err
	}

	return nil
}
