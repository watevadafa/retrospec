package utils

import (
	"regexp"
	"strings"

	"github.com/watevadafa/retrospec/internal/types"
	"golang.org/x/crypto/bcrypt"
)

func IsNullOrEmpty(value *string) bool {
	return value == nil || strings.TrimSpace(*value) == ""
}

func IsValidEmail(value *string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return emailRegex.MatchString(strings.ToLower(*value))
}

func IsValidPassword(value *string, personalInfo ...types.PersonalInfo) bool {
	// Password validation
	if len(*value) < 8 || len(*value) > 64 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(*value)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(*value)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(*value)
	hasSymbol := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>\-]`).MatchString(*value)

	if !hasUpper || !hasLower || !hasNumber || !hasSymbol {
		return false
	}

	// If personal info is provided, check against it
	if len(personalInfo) > 0 {
		passwordLower := strings.ToLower(*value)
		info := personalInfo[0]

		// Check if password contains parts of personal info
		personalInfo := []string{
			strings.ToLower(info.FirstName),
			strings.ToLower(*info.LastName),
			strings.Split(strings.ToLower(info.Email), "@")[0],
		}

		for _, info := range personalInfo {
			if info != "" && len(info) > 2 && strings.Contains(passwordLower, info) {
				return false
			}
		}
	}

	return true
}

func HashPassword(password *string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
