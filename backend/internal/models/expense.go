package models

import (
	"time"

	"github.com/google/uuid"
)

// ExpenseCategory represents an expense category
type ExpenseCategory struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description,omitempty" db:"description"`
	Color       string    `json:"color" db:"color"`
	Icon        *string   `json:"icon,omitempty" db:"icon"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Expense represents an expense entry
type Expense struct {
	ID            uuid.UUID `json:"id" db:"id"`
	UserID        uuid.UUID `json:"user_id" db:"user_id"`
	CategoryID    uuid.UUID `json:"category_id" db:"category_id"`
	Amount        float64   `json:"amount" db:"amount"`
	Description   string    `json:"description" db:"description"`
	ExpenseDate   time.Time `json:"expense_date" db:"expense_date"`
	PaymentMethod *string   `json:"payment_method,omitempty" db:"payment_method"`
	Location      *string   `json:"location,omitempty" db:"location"`
	ReceiptURL    *string   `json:"receipt_url,omitempty" db:"receipt_url"`
	Tags          []string  `json:"tags,omitempty" db:"tags"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`

	// Relations
	Category *ExpenseCategory `json:"category,omitempty"`
	User     *User            `json:"user,omitempty"`
}

// ExpenseCreateRequest represents the request to create a new expense
type ExpenseCreateRequest struct {
	CategoryID    uuid.UUID `json:"category_id" validate:"required"`
	Amount        float64   `json:"amount" validate:"required,gt=0"`
	Description   string    `json:"description" validate:"required"`
	ExpenseDate   time.Time `json:"expense_date" validate:"required"`
	PaymentMethod *string   `json:"payment_method,omitempty"`
	Location      *string   `json:"location,omitempty"`
	ReceiptURL    *string   `json:"receipt_url,omitempty"`
	Tags          []string  `json:"tags,omitempty"`
}

// ExpenseUpdateRequest represents the request to update an expense
type ExpenseUpdateRequest struct {
	CategoryID    *uuid.UUID `json:"category_id,omitempty"`
	Amount        *float64   `json:"amount,omitempty" validate:"omitempty,gt=0"`
	Description   *string    `json:"description,omitempty"`
	ExpenseDate   *time.Time `json:"expense_date,omitempty"`
	PaymentMethod *string    `json:"payment_method,omitempty"`
	Location      *string    `json:"location,omitempty"`
	ReceiptURL    *string    `json:"receipt_url,omitempty"`
	Tags          []string   `json:"tags,omitempty"`
}

// ExpenseFilter represents filters for expense queries
type ExpenseFilter struct {
	UserID        uuid.UUID  `json:"user_id"`
	CategoryID    *uuid.UUID `json:"category_id,omitempty"`
	StartDate     *time.Time `json:"start_date,omitempty"`
	EndDate       *time.Time `json:"end_date,omitempty"`
	MinAmount     *float64   `json:"min_amount,omitempty"`
	MaxAmount     *float64   `json:"max_amount,omitempty"`
	PaymentMethod *string    `json:"payment_method,omitempty"`
	Tags          []string   `json:"tags,omitempty"`
	Limit         int        `json:"limit,omitempty"`
	Offset        int        `json:"offset,omitempty"`
}

// ExpenseSummary represents expense summary statistics
type ExpenseSummary struct {
	TotalAmount     float64                  `json:"total_amount"`
	TotalCount      int                      `json:"total_count"`
	AverageAmount   float64                  `json:"average_amount"`
	ByCategory      []CategoryExpenseSummary `json:"by_category,omitempty"`
	ByMonth         []MonthlyExpenseSummary  `json:"by_month,omitempty"`
	ByPaymentMethod []PaymentMethodSummary   `json:"by_payment_method,omitempty"`
}

// CategoryExpenseSummary represents expense summary by category
type CategoryExpenseSummary struct {
	CategoryID   uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Amount       float64   `json:"amount"`
	Count        int       `json:"count"`
	Percentage   float64   `json:"percentage"`
}

// MonthlyExpenseSummary represents expense summary by month
type MonthlyExpenseSummary struct {
	Year   int     `json:"year"`
	Month  int     `json:"month"`
	Amount float64 `json:"amount"`
	Count  int     `json:"count"`
}

// PaymentMethodSummary represents expense summary by payment method
type PaymentMethodSummary struct {
	PaymentMethod string  `json:"payment_method"`
	Amount        float64 `json:"amount"`
	Count         int     `json:"count"`
	Percentage    float64 `json:"percentage"`
}
