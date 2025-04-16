package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPass, plainPass string) bool {
	byteHash := []byte(hashedPass)
	password := []byte(plainPass)

	err := bcrypt.CompareHashAndPassword(byteHash, password)
	return err == nil
}
