package validator

import (
	"fmt"
	"regexp"
)

func CheckEmpty(data ...interface{})error{
	for _, v := range data {
		if v == nil || v == "" {
			return fmt.Errorf("field cannot be empty")
		}
	}
	return nil
}

func CheckEmail(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	isValidEmail := regexp.MustCompile(emailRegex).MatchString
	// Simple regex for email validation
	if !isValidEmail(email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

func CheckPassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	hasLetter := regexp.MustCompile(`[A-Za-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)

	// Password must be at least 8 characters long and contain at least one letter and one number
	if !hasLetter || !hasDigit {
		return fmt.Errorf("password must contain at least one letter and one number")
	}
	return nil
}