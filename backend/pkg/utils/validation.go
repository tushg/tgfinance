package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return ""
	}

	var messages []string
	for _, err := range e {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "; ")
}

// HasErrors returns true if there are validation errors
func (e ValidationErrors) HasErrors() bool {
	return len(e) > 0
}

// Add adds a validation error
func (e *ValidationErrors) Add(field, message string) {
	*e = append(*e, ValidationError{Field: field, Message: message})
}

// Email validation regex pattern
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// ValidateEmail validates email format
func ValidateEmail(email string) error {
	if email == "" {
		return &ValidationError{Field: "email", Message: "email is required"}
	}

	if !emailRegex.MatchString(email) {
		return &ValidationError{Field: "email", Message: "invalid email format"}
	}

	if len(email) > 254 {
		return &ValidationError{Field: "email", Message: "email too long (max 254 characters)"}
	}

	return nil
}

// ValidatePassword validates password strength
func ValidatePassword(password string) error {
	if password == "" {
		return &ValidationError{Field: "password", Message: "password is required"}
	}

	if len(password) < 8 {
		return &ValidationError{Field: "password", Message: "password must be at least 8 characters long"}
	}

	if len(password) > 128 {
		return &ValidationError{Field: "password", Message: "password too long (max 128 characters)"}
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
		return &ValidationError{Field: "password", Message: "password must contain at least one uppercase letter"}
	}

	if !hasLower {
		return &ValidationError{Field: "password", Message: "password must contain at least one lowercase letter"}
	}

	if !hasNumber {
		return &ValidationError{Field: "password", Message: "password must contain at least one number"}
	}

	if !hasSpecial {
		return &ValidationError{Field: "password", Message: "password must contain at least one special character"}
	}

	return nil
}

// ValidateRequired validates that a field is not empty
func ValidateRequired(value, fieldName string) error {
	if strings.TrimSpace(value) == "" {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s is required", fieldName)}
	}
	return nil
}

// ValidateLength validates string length
func ValidateLength(value, fieldName string, min, max int) error {
	length := len(strings.TrimSpace(value))

	if min > 0 && length < min {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s must be at least %d characters long", fieldName, min)}
	}

	if max > 0 && length > max {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s must be no more than %d characters long", fieldName, max)}
	}

	return nil
}

// ValidateName validates name format (letters, spaces, hyphens, apostrophes)
func ValidateName(name, fieldName string) error {
	if err := ValidateRequired(name, fieldName); err != nil {
		return err
	}

	if err := ValidateLength(name, fieldName, 2, 100); err != nil {
		return err
	}

	// Allow letters, spaces, hyphens, and apostrophes
	nameRegex := regexp.MustCompile(`^[a-zA-Z\s\-']+$`)
	if !nameRegex.MatchString(name) {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s can only contain letters, spaces, hyphens, and apostrophes", fieldName)}
	}

	return nil
}

// ValidatePhone validates phone number format
func ValidatePhone(phone string) error {
	if phone == "" {
		return &ValidationError{Field: "phone", Message: "phone number is required"}
	}

	// Remove all non-digit characters
	digits := regexp.MustCompile(`\D`).ReplaceAllString(phone, "")

	if len(digits) < 10 || len(digits) > 15 {
		return &ValidationError{Field: "phone", Message: "phone number must be between 10 and 15 digits"}
	}

	return nil
}

// ValidateAmount validates monetary amount
func ValidateAmount(amount float64, fieldName string) error {
	if amount <= 0 {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s must be greater than 0", fieldName)}
	}

	if amount > 999999999.99 {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s is too large (max 999,999,999.99)", fieldName)}
	}

	return nil
}

// ValidateDate validates date format (YYYY-MM-DD)
func ValidateDate(date, fieldName string) error {
	if err := ValidateRequired(date, fieldName); err != nil {
		return err
	}

	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !dateRegex.MatchString(date) {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s must be in YYYY-MM-DD format", fieldName)}
	}

	// Parse the date to validate it's actually valid
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s must be a valid date", fieldName)}
	}

	return nil
}

// ValidateUUID validates UUID format
func ValidateUUID(uuid, fieldName string) error {
	if err := ValidateRequired(uuid, fieldName); err != nil {
		return err
	}

	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	if !uuidRegex.MatchString(strings.ToLower(uuid)) {
		return &ValidationError{Field: fieldName, Message: fmt.Sprintf("%s must be a valid UUID", fieldName)}
	}

	return nil
}

// ValidatePagination validates pagination parameters
func ValidatePagination(page, limit int) error {
	if page < 1 {
		return &ValidationError{Field: "page", Message: "page must be greater than 0"}
	}

	if limit < 1 || limit > 100 {
		return &ValidationError{Field: "limit", Message: "limit must be between 1 and 100"}
	}

	return nil
}

// ValidateSortOrder validates sort order parameter
func ValidateSortOrder(sortOrder string) error {
	if sortOrder == "" {
		return nil // Default sort order is allowed
	}

	sortOrder = strings.ToLower(sortOrder)
	if sortOrder != "asc" && sortOrder != "desc" {
		return &ValidationError{Field: "sort_order", Message: "sort_order must be 'asc' or 'desc'"}
	}

	return nil
}
