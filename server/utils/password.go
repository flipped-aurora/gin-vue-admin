package utils

import (
	"crypto/rand"
	"errors"
)

const passwordAlphabet = "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz23456789+-@!"

func RandomPassword(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("password length must be greater than 0")
	}

	buf := make([]byte, length)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}

	password := make([]byte, length)
	for i, b := range buf {
		password[i] = passwordAlphabet[int(b)%len(passwordAlphabet)]
	}

	return string(password), nil
}
