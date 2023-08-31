package validator

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	_, err := regexp.MatchString(emailPattern, email)
	if err != nil {
		return errors.New("INVALID EMAIL ADDRESS")
	}
	return nil
}

func ValidatePhone(phone string) error {
	phonePattern := `^(\+62|62|0)8[1-9][0-9]{12,13}$`
	_, err := regexp.MatchString(phonePattern, phone)
	if err != nil {
		return errors.New("INVALID PHONE NUMBER")
	}
	return nil
}
