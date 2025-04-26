package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user details
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Scheme represents a scholarship scheme
type Scheme struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name" gorm:"not null"`
	Description     string         `json:"description"`
	Eligibility     string         `json:"eligibility" gorm:"type:text"`                      // who can apply
	Amount          float64        `json:"amount"`                                            // scholarship amount
	ApplicationLink string         `json:"application_link"`                                  // where to apply
	StartDate       time.Time      `json:"start_date"`                                        // when applications open
	EndDate         time.Time      `json:"end_date"`                                          // when applications close
	Status          string         `json:"status" gorm:"type:varchar(20);default:'upcoming'"` // open, closed, upcoming
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}

// SchemeFilter represents the filter criteria for fetching schemes
type SchemeFilter struct {
	Name        *string    `form:"name"`        // partial or exact match
	Status      *string    `form:"status"`      // open, closed, upcoming
	MinAmount   *float64   `form:"min_amount"`  // minimum scholarship amount
	MaxAmount   *float64   `form:"max_amount"`  // maximum scholarship amount
	StartAfter  *time.Time `form:"start_after"` // start date after
	EndBefore   *time.Time `form:"end_before"`  // end date before
	Eligibility *string    `form:"eligibility"` // keyword in eligibility
}

// Application represents a submitted application for a scheme
type Application struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"` // Link to User
	SchemeID  uint      `gorm:"not null" json:"scheme_id"`
	Status    string    `gorm:"type:varchar(20);default:'pending'" json:"status"` // pending, approved, rejected
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

// ErrorResponse represents a structured error response
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// SuccessResponse represents a structured success response
type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type SchemeResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    interface{}     `json:"data,omitempty"`
	Meta    *PaginationMeta `json:"meta,omitempty"`
}

type PaginationMeta struct {
	ResourceCount int    `json:"resource_count" example:"200"`
	TotalPages    int64  `json:"total_pages,omitempty" example:"20"`
	Page          int64  `json:"page,omitempty" example:"10"`
	Limit         int64  `json:"limit,omitempty" example:"10"`
	Next          string `json:"next,omitempty" example:"/api/v1/schemes?limit=10&page=11"`
	Previous      string `json:"previous,omitempty" example:"/api/v1/schemes?limit=10&page=9"`
}

// The PaginationInput struct represents the input required for pagination.
type PaginationInput struct {
	Page  int64 `json:"page" example:"10"`
	Limit int64 `json:"limit" example:"10"`
}

// PaginationParse interface processes the pagination input.
type PaginationParse interface {
	GetOffset() int64
	GetLimit() int64
}

// GetOffset returns the offset value for gorm.
func (p PaginationInput) GetOffset() int64 {
	return (p.Page - 1) * p.Limit
}

// GetLimit returns the limit value for gorm.
func (p PaginationInput) GetLimit() int64 {
	return p.Limit
}
