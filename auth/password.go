package auth

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(input string) (string, error) {
	buf := bytes.NewBufferString(input)
	if buf == nil {
		return "", fmt.Errorf("input string to buffer is nil")
	}
	if buf.Len() > 72 {
		return "", fmt.Errorf("input string is longer than 72 bytes")
	}

	hash, err := bcrypt.GenerateFromPassword(buf.Bytes(), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyPassword(input string, hash string) error {
	inputBuf := bytes.NewBufferString(input)
	if inputBuf == nil {
		return fmt.Errorf("input string to buffer is nil")
	}

	hashBuf := bytes.NewBufferString(hash)
	if hashBuf == nil {
		return fmt.Errorf("hash string to buffer is nil")
	}

	if err := bcrypt.CompareHashAndPassword(hashBuf.Bytes(), inputBuf.Bytes()); err != nil {
		return err
	}

	return nil
}
