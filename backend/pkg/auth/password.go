package auth

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// PasswordManager handles password hashing and verification
type PasswordManager struct {
	cost int
}

// NewPasswordManager creates a new password manager
func NewPasswordManager() *PasswordManager {
	return &PasswordManager{
		cost: bcrypt.DefaultCost, // 10 rounds
	}
}

// HashPassword hashes a password using bcrypt
func (pm *PasswordManager) HashPassword(password string) (string, error) {
	// Validate password strength before hashing
	if err := pm.validatePasswordStrength(password); err != nil {
		return "", err
	}

	// Hash the password
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), pm.cost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

// VerifyPassword verifies a password against its hash
func (pm *PasswordManager) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// validatePasswordStrength validates password requirements
func (pm *PasswordManager) validatePasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	if len(password) > 128 {
		return errors.New("password must be less than 128 characters")
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}

	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}

	if !hasNumber {
		return errors.New("password must contain at least one number")
	}

	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

// IsPasswordValid checks if a password meets strength requirements without hashing
func (pm *PasswordManager) IsPasswordValid(password string) error {
	return pm.validatePasswordStrength(password)
}

// GetPasswordStrength returns a strength score (0-100) for a password
func (pm *PasswordManager) GetPasswordStrength(password string) int {
	score := 0

	// Length contribution
	if len(password) >= 8 {
		score += 20
	}
	if len(password) >= 12 {
		score += 10
	}
	if len(password) >= 16 {
		score += 10
	}

	// Character variety contribution
	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if hasUpper {
		score += 15
	}
	if hasLower {
		score += 15
	}
	if hasNumber {
		score += 15
	}
	if hasSpecial {
		score += 15
	}

	// Bonus for mixed case
	if hasUpper && hasLower {
		score += 10
	}

	return score
}

// GetPasswordStrengthLabel returns a human-readable strength label
func (pm *PasswordManager) GetPasswordStrengthLabel(password string) string {
	score := pm.GetPasswordStrength(password)

	switch {
	case score >= 80:
		return "Very Strong"
	case score >= 60:
		return "Strong"
	case score >= 40:
		return "Medium"
	case score >= 20:
		return "Weak"
	default:
		return "Very Weak"
	}
}
