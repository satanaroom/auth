package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(hashedPassword, candidatePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
	return err == nil
}

func GeneratePasswordHash(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("generate from password: %w", err)
	}

	return string(hashedBytes[:]), nil
}
