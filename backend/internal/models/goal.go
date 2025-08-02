package models

import (
	"time"

	"github.com/google/uuid"
)

// FinancialGoal represents a financial goal
type FinancialGoal struct {
	ID            uuid.UUID  `json:"id" db:"id"`
	UserID        uuid.UUID  `json:"user_id" db:"user_id"`
	Name          string     `json:"name" db:"name"`
	Description   *string    `json:"description,omitempty" db:"description"`
	TargetAmount  float64    `json:"target_amount" db:"target_amount"`
	CurrentAmount float64    `json:"current_amount" db:"current_amount"`
	TargetDate    *time.Time `json:"target_date,omitempty" db:"target_date"`
	GoalType      string     `json:"goal_type" db:"goal_type"`
	Priority      string     `json:"priority" db:"priority"`
	Status        string     `json:"status" db:"status"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`

	// Relations
	User *User `json:"user,omitempty"`
}

// GoalContribution represents a contribution to a financial goal
type GoalContribution struct {
	ID               uuid.UUID `json:"id" db:"id"`
	GoalID           uuid.UUID `json:"goal_id" db:"goal_id"`
	Amount           float64   `json:"amount" db:"amount"`
	ContributionDate time.Time `json:"contribution_date" db:"contribution_date"`
	Source           *string   `json:"source,omitempty" db:"source"`
	Notes            *string   `json:"notes,omitempty" db:"notes"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`

	// Relations
	Goal *FinancialGoal `json:"goal,omitempty"`
}

// GoalCreateRequest represents the request to create a new financial goal
type GoalCreateRequest struct {
	Name         string     `json:"name" validate:"required"`
	Description  *string    `json:"description,omitempty"`
	TargetAmount float64    `json:"target_amount" validate:"required,gt=0"`
	TargetDate   *time.Time `json:"target_date,omitempty"`
	GoalType     string     `json:"goal_type" validate:"required,oneof=savings investment debt_payoff purchase emergency_fund"`
	Priority     string     `json:"priority" validate:"required,oneof=low medium high"`
}

// GoalUpdateRequest represents the request to update a financial goal
type GoalUpdateRequest struct {
	Name         *string    `json:"name,omitempty"`
	Description  *string    `json:"description,omitempty"`
	TargetAmount *float64   `json:"target_amount,omitempty" validate:"omitempty,gt=0"`
	TargetDate   *time.Time `json:"target_date,omitempty"`
	GoalType     *string    `json:"goal_type,omitempty"`
	Priority     *string    `json:"priority,omitempty"`
	Status       *string    `json:"status,omitempty"`
}

// GoalContributionCreateRequest represents the request to create a goal contribution
type GoalContributionCreateRequest struct {
	Amount           float64   `json:"amount" validate:"required,gt=0"`
	ContributionDate time.Time `json:"contribution_date" validate:"required"`
	Source           *string   `json:"source,omitempty"`
	Notes            *string   `json:"notes,omitempty"`
}

// GoalFilter represents filters for goal queries
type GoalFilter struct {
	UserID   uuid.UUID `json:"user_id"`
	GoalType *string   `json:"goal_type,omitempty"`
	Priority *string   `json:"priority,omitempty"`
	Status   *string   `json:"status,omitempty"`
	Limit    int       `json:"limit,omitempty"`
	Offset   int       `json:"offset,omitempty"`
}

// GoalSummary represents goal summary statistics
type GoalSummary struct {
	TotalGoals         int                   `json:"total_goals"`
	ActiveGoals        int                   `json:"active_goals"`
	CompletedGoals     int                   `json:"completed_goals"`
	TotalTargetAmount  float64               `json:"total_target_amount"`
	TotalCurrentAmount float64               `json:"total_current_amount"`
	TotalProgress      float64               `json:"total_progress"`
	ByType             []TypeGoalSummary     `json:"by_type,omitempty"`
	ByPriority         []PriorityGoalSummary `json:"by_priority,omitempty"`
	ByStatus           []StatusGoalSummary   `json:"by_status,omitempty"`
}

// TypeGoalSummary represents goal summary by type
type TypeGoalSummary struct {
	GoalType      string  `json:"goal_type"`
	Count         int     `json:"count"`
	TargetAmount  float64 `json:"target_amount"`
	CurrentAmount float64 `json:"current_amount"`
	Progress      float64 `json:"progress"`
}

// PriorityGoalSummary represents goal summary by priority
type PriorityGoalSummary struct {
	Priority      string  `json:"priority"`
	Count         int     `json:"count"`
	TargetAmount  float64 `json:"target_amount"`
	CurrentAmount float64 `json:"current_amount"`
	Progress      float64 `json:"progress"`
}

// StatusGoalSummary represents goal summary by status
type StatusGoalSummary struct {
	Status        string  `json:"status"`
	Count         int     `json:"count"`
	TargetAmount  float64 `json:"target_amount"`
	CurrentAmount float64 `json:"current_amount"`
	Progress      float64 `json:"progress"`
}

// GetProgress returns the progress percentage of the goal
func (g *FinancialGoal) GetProgress() float64 {
	if g.TargetAmount == 0 {
		return 0
	}
	return (g.CurrentAmount / g.TargetAmount) * 100
}

// IsCompleted returns true if the goal is completed
func (g *FinancialGoal) IsCompleted() bool {
	return g.CurrentAmount >= g.TargetAmount
}

// IsOverdue returns true if the goal is overdue
func (g *FinancialGoal) IsOverdue() bool {
	if g.TargetDate == nil {
		return false
	}
	return time.Now().After(*g.TargetDate) && !g.IsCompleted()
}
