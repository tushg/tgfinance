package auth

import (
	"testing"

	"github.com/google/uuid"
)

func TestJWTManager(t *testing.T) {
	jwtManager := NewJWTManager()
	userID := uuid.New()
	email := "test@example.com"

	// Test token generation
	token, err := jwtManager.GenerateToken(userID, email)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Test token validation
	claims, err := jwtManager.ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected user ID %v, got %v", userID, claims.UserID)
	}

	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}

	// Test refresh token
	refreshToken, err := jwtManager.GenerateRefreshToken(userID)
	if err != nil {
		t.Fatalf("Failed to generate refresh token: %v", err)
	}

	refreshClaims, err := jwtManager.ValidateToken(refreshToken)
	if err != nil {
		t.Fatalf("Failed to validate refresh token: %v", err)
	}

	if refreshClaims.UserID != userID {
		t.Errorf("Expected user ID %v, got %v", userID, refreshClaims.UserID)
	}
}

func TestPasswordManager(t *testing.T) {
	passwordManager := NewPasswordManager()

	// Test password validation
	validPassword := "SecurePass123!"
	err := passwordManager.IsPasswordValid(validPassword)
	if err != nil {
		t.Fatalf("Valid password should pass validation: %v", err)
	}

	// Test weak password
	weakPassword := "weak"
	err = passwordManager.IsPasswordValid(weakPassword)
	if err == nil {
		t.Error("Weak password should fail validation")
	}

	// Test password hashing
	password := "SecurePass123!"
	hashedPassword, err := passwordManager.HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	// Test password verification
	err = passwordManager.VerifyPassword(hashedPassword, password)
	if err != nil {
		t.Fatalf("Failed to verify password: %v", err)
	}

	// Test wrong password
	err = passwordManager.VerifyPassword(hashedPassword, "wrongpassword")
	if err == nil {
		t.Error("Wrong password should fail verification")
	}

	// Test password strength
	strength := passwordManager.GetPasswordStrength(password)
	if strength < 60 {
		t.Errorf("Password should be strong, got strength: %d", strength)
	}

	label := passwordManager.GetPasswordStrengthLabel(password)
	if label != "Strong" && label != "Very Strong" {
		t.Errorf("Expected strong password label, got: %s", label)
	}
}

func TestPasswordStrength(t *testing.T) {
	passwordManager := NewPasswordManager()

	testCases := []struct {
		password string
		expected string
	}{
		{"Abc123!@#", "Very Strong"},
		{"weak", "Very Weak"},
		{"password123", "Medium"},
		{"Password123", "Strong"},
		{"SuperSecurePass123!@#", "Very Strong"},
	}

	for _, tc := range testCases {
		label := passwordManager.GetPasswordStrengthLabel(tc.password)
		if label != tc.expected {
			t.Errorf("Password '%s' expected %s, got %s", tc.password, tc.expected, label)
		}
	}
}
