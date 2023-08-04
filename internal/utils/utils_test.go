package utils

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestGeneratePasswordHash(t *testing.T) {
	password := "mysecretpassword"
	hash, err := GeneratePasswordHash(password)

	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Errorf("Hashed password does not match original: %v", err)
	}
}

func TestComparePasswords(t *testing.T) {
	password := "mysecretpassword"
	hashed, err := GeneratePasswordHash(password)
	if err != nil {
		t.Fatalf("Error hashing password: %v", err)
	}

	if !HashPassword(hashed, password) {
		t.Errorf("Expected passwords to match, but they did not.")
	}

	wrongPassword := "incorrectpassword"
	if HashPassword(hashed, wrongPassword) {
		t.Errorf("Expected passwords not to match, but they did.")
	}
}
