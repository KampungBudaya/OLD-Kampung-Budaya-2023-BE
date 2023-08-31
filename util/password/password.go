package password

import (
	"errors"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) (string, error) {
	saltRound, err := strconv.Atoi(os.Getenv("SALT_ROUND"))
	if err != nil {
		return "", errors.New("FAILED TO PARSING STRING INTO NUMBER")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), saltRound)
	if err != nil {
		return "", errors.New("FAILED TO HASHING PASSWORD")
	}

	return string(hashedPassword), nil
}

func ComparePassword(password, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return errors.New("WRONG PASSWORD")
	}
	return nil
}
