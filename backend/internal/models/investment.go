package models

import (
	"time"

	"github.com/google/uuid"
)

// InvestmentType represents an investment type
type InvestmentType struct {
	ID             uuid.UUID `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	Description    *string   `json:"description,omitempty" db:"description"`
	RiskLevel      string    `json:"risk_level" db:"risk_level"`
	ExpectedReturn float64   `json:"expected_return" db:"expected_return"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Investment represents an investment entry
type Investment struct {
	ID            uuid.UUID  `json:"id" db:"id"`
	UserID        uuid.UUID  `json:"user_id" db:"user_id"`
	TypeID        uuid.UUID  `json:"type_id" db:"type_id"`
	Name          string     `json:"name" db:"name"`
	Amount        float64    `json:"amount" db:"amount"`
	CurrentValue  *float64   `json:"current_value,omitempty" db:"current_value"`
	StartDate     time.Time  `json:"start_date" db:"start_date"`
	EndDate       *time.Time `json:"end_date,omitempty" db:"end_date"`
	InterestRate  *float64   `json:"interest_rate,omitempty" db:"interest_rate"`
	Institution   *string    `json:"institution,omitempty" db:"institution"`
	AccountNumber *string    `json:"account_number,omitempty" db:"account_number"`
	Notes         *string    `json:"notes,omitempty" db:"notes"`
	Status        string     `json:"status" db:"status"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`

	// Relations
	Type *InvestmentType `json:"type,omitempty"`
	User *User           `json:"user,omitempty"`
}

// InvestmentTransaction represents an investment transaction
type InvestmentTransaction struct {
	ID              uuid.UUID `json:"id" db:"id"`
	InvestmentID    uuid.UUID `json:"investment_id" db:"investment_id"`
	TransactionType string    `json:"transaction_type" db:"transaction_type"`
	Amount          float64   `json:"amount" db:"amount"`
	TransactionDate time.Time `json:"transaction_date" db:"transaction_date"`
	Description     *string   `json:"description,omitempty" db:"description"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`

	// Relations
	Investment *Investment `json:"investment,omitempty"`
}

// InvestmentCreateRequest represents the request to create a new investment
type InvestmentCreateRequest struct {
	TypeID        uuid.UUID  `json:"type_id" validate:"required"`
	Name          string     `json:"name" validate:"required"`
	Amount        float64    `json:"amount" validate:"required,gt=0"`
	CurrentValue  *float64   `json:"current_value,omitempty"`
	StartDate     time.Time  `json:"start_date" validate:"required"`
	EndDate       *time.Time `json:"end_date,omitempty"`
	InterestRate  *float64   `json:"interest_rate,omitempty"`
	Institution   *string    `json:"institution,omitempty"`
	AccountNumber *string    `json:"account_number,omitempty"`
	Notes         *string    `json:"notes,omitempty"`
}

// InvestmentUpdateRequest represents the request to update an investment
type InvestmentUpdateRequest struct {
	Name          *string    `json:"name,omitempty"`
	Amount        *float64   `json:"amount,omitempty" validate:"omitempty,gt=0"`
	CurrentValue  *float64   `json:"current_value,omitempty"`
	EndDate       *time.Time `json:"end_date,omitempty"`
	InterestRate  *float64   `json:"interest_rate,omitempty"`
	Institution   *string    `json:"institution,omitempty"`
	AccountNumber *string    `json:"account_number,omitempty"`
	Notes         *string    `json:"notes,omitempty"`
	Status        *string    `json:"status,omitempty"`
}

// InvestmentTransactionCreateRequest represents the request to create a transaction
type InvestmentTransactionCreateRequest struct {
	TransactionType string    `json:"transaction_type" validate:"required,oneof=deposit withdrawal interest dividend"`
	Amount          float64   `json:"amount" validate:"required,gt=0"`
	TransactionDate time.Time `json:"transaction_date" validate:"required"`
	Description     *string   `json:"description,omitempty"`
}

// InvestmentFilter represents filters for investment queries
type InvestmentFilter struct {
	UserID      uuid.UUID  `json:"user_id"`
	TypeID      *uuid.UUID `json:"type_id,omitempty"`
	Status      *string    `json:"status,omitempty"`
	Institution *string    `json:"institution,omitempty"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	Limit       int        `json:"limit,omitempty"`
	Offset      int        `json:"offset,omitempty"`
}

// InvestmentSummary represents investment summary statistics
type InvestmentSummary struct {
	TotalInvested     float64                   `json:"total_invested"`
	TotalCurrentValue float64                   `json:"total_current_value"`
	TotalGain         float64                   `json:"total_gain"`
	TotalGainPercent  float64                   `json:"total_gain_percent"`
	ByType            []TypeInvestmentSummary   `json:"by_type,omitempty"`
	ByStatus          []StatusInvestmentSummary `json:"by_status,omitempty"`
	ByInstitution     []InstitutionSummary      `json:"by_institution,omitempty"`
}

// TypeInvestmentSummary represents investment summary by type
type TypeInvestmentSummary struct {
	TypeID         uuid.UUID `json:"type_id"`
	TypeName       string    `json:"type_name"`
	InvestedAmount float64   `json:"invested_amount"`
	CurrentValue   float64   `json:"current_value"`
	Gain           float64   `json:"gain"`
	GainPercent    float64   `json:"gain_percent"`
	Count          int       `json:"count"`
}

// StatusInvestmentSummary represents investment summary by status
type StatusInvestmentSummary struct {
	Status         string  `json:"status"`
	InvestedAmount float64 `json:"invested_amount"`
	CurrentValue   float64 `json:"current_value"`
	Gain           float64 `json:"gain"`
	Count          int     `json:"count"`
}

// InstitutionSummary represents investment summary by institution
type InstitutionSummary struct {
	Institution    string  `json:"institution"`
	InvestedAmount float64 `json:"invested_amount"`
	CurrentValue   float64 `json:"current_value"`
	Gain           float64 `json:"gain"`
	Count          int     `json:"count"`
}
