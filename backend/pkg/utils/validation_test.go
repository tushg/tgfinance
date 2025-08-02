package utils

import (
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{"valid email", "test@example.com", false},
		{"valid email with subdomain", "test@sub.example.com", false},
		{"valid email with plus", "test+tag@example.com", false},
		{"valid email with dots", "test.name@example.com", false},
		{"empty email", "", true},
		{"invalid format", "invalid-email", true},
		{"missing domain", "test@", true},
		{"missing @", "testexample.com", true},
		{"too long", "a" + string(make([]byte, 254)) + "@example.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"valid password", "StrongPass123!", false},
		{"valid password with symbols", "MyP@ssw0rd#", false},
		{"empty password", "", true},
		{"too short", "Abc123!", true},
		{"too long", string(make([]byte, 129)), true},
		{"no uppercase", "password123!", true},
		{"no lowercase", "PASSWORD123!", true},
		{"no number", "Password!", true},
		{"no special char", "Password123", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateRequired(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		fieldName string
		wantErr   bool
	}{
		{"valid value", "test", "field", false},
		{"empty string", "", "field", true},
		{"whitespace only", "   ", "field", true},
		{"tab only", "\t", "field", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRequired(tt.value, tt.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateRequired() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateLength(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		fieldName string
		min       int
		max       int
		wantErr   bool
	}{
		{"valid length", "test", "field", 2, 10, false},
		{"exact min length", "ab", "field", 2, 10, false},
		{"exact max length", "abcdefghij", "field", 2, 10, false},
		{"too short", "a", "field", 2, 10, true},
		{"too long", "abcdefghijk", "field", 2, 10, true},
		{"empty with min 0", "", "field", 0, 10, false},
		{"whitespace with min 0", "   ", "field", 0, 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateLength(tt.value, tt.fieldName, tt.min, tt.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateName(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		fieldName string
		wantErr   bool
	}{
		{"valid name", "John Doe", "name", false},
		{"valid name with hyphen", "Jean-Pierre", "name", false},
		{"valid name with apostrophe", "O'Connor", "name", false},
		{"single character", "J", "name", true},
		{"too long", string(make([]byte, 101)), "name", true},
		{"contains numbers", "John123", "name", true},
		{"contains special chars", "John@Doe", "name", true},
		{"empty", "", "name", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateName(tt.value, tt.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePhone(t *testing.T) {
	tests := []struct {
		name    string
		phone   string
		wantErr bool
	}{
		{"valid phone", "1234567890", false},
		{"valid phone with formatting", "(123) 456-7890", false},
		{"valid international", "+1-234-567-8901", false},
		{"empty phone", "", true},
		{"too short", "123456789", true},
		{"too long", "1234567890123456", true},
		{"non-numeric", "abcdefghij", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePhone(tt.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePhone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAmount(t *testing.T) {
	tests := []struct {
		name      string
		amount    float64
		fieldName string
		wantErr   bool
	}{
		{"valid amount", 100.50, "amount", false},
		{"valid small amount", 0.01, "amount", false},
		{"valid large amount", 999999999.99, "amount", false},
		{"zero amount", 0, "amount", true},
		{"negative amount", -100, "amount", true},
		{"too large", 1000000000, "amount", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAmount(tt.amount, tt.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateDate(t *testing.T) {
	tests := []struct {
		name      string
		date      string
		fieldName string
		wantErr   bool
	}{
		{"valid date", "2023-12-25", "date", false},
		{"valid date with leading zeros", "2023-01-01", "date", false},
		{"empty date", "", "date", true},
		{"invalid format", "2023/12/25", "date", true},
		{"invalid format 2", "25-12-2023", "date", true},
		{"invalid format 3", "2023-12-25T10:30:00", "date", true},
		{"invalid date", "2023-13-45", "date", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDate(tt.date, tt.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateUUID(t *testing.T) {
	tests := []struct {
		name      string
		uuid      string
		fieldName string
		wantErr   bool
	}{
		{"valid UUID", "550e8400-e29b-41d4-a716-446655440000", "id", false},
		{"valid UUID uppercase", "550E8400-E29B-41D4-A716-446655440000", "id", false},
		{"empty UUID", "", "id", true},
		{"invalid format", "550e8400-e29b-41d4-a716-44665544000", "id", true},
		{"invalid format 2", "550e8400-e29b-41d4-a716-4466554400000", "id", true},
		{"invalid characters", "550e8400-e29b-41d4-a716-44665544000g", "id", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUUID(tt.uuid, tt.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateUUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePagination(t *testing.T) {
	tests := []struct {
		name    string
		page    int
		limit   int
		wantErr bool
	}{
		{"valid pagination", 1, 10, false},
		{"valid pagination 2", 5, 50, false},
		{"valid pagination 3", 10, 100, false},
		{"zero page", 0, 10, true},
		{"negative page", -1, 10, true},
		{"zero limit", 1, 0, true},
		{"negative limit", 1, -1, true},
		{"limit too high", 1, 101, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePagination(tt.page, tt.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePagination() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateSortOrder(t *testing.T) {
	tests := []struct {
		name      string
		sortOrder string
		wantErr   bool
	}{
		{"valid asc", "asc", false},
		{"valid desc", "desc", false},
		{"valid ASC", "ASC", false},
		{"valid DESC", "DESC", false},
		{"empty sort order", "", false},
		{"invalid sort order", "invalid", true},
		{"invalid sort order 2", "ascending", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSortOrder(tt.sortOrder)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSortOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidationErrors(t *testing.T) {
	var errors ValidationErrors

	// Test empty errors
	if errors.HasErrors() {
		t.Error("Empty ValidationErrors should not have errors")
	}

	if errors.Error() != "" {
		t.Error("Empty ValidationErrors should return empty string")
	}

	// Test adding errors
	errors.Add("field1", "error1")
	errors.Add("field2", "error2")

	if !errors.HasErrors() {
		t.Error("ValidationErrors with errors should have errors")
	}

	expected := "field1: error1; field2: error2"
	if errors.Error() != expected {
		t.Errorf("ValidationErrors.Error() = %v, want %v", errors.Error(), expected)
	}
}
