package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Note struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	Title     string    `gorm:"type:varchar(255);uniqueIndex:idx_notes_title,LENGTH(255);not null" json:"title,omitempty"`
	Content   string    `gorm:"not null" json:"content,omitempty"`
	Category  string    `gorm:"varchar(100)" json:"category,omitempty"`
	Published bool      `gorm:"default:false;not null" json:"published"`
	CreatedAt time.Time `gorm:"not null;default:'1970-01-01 00:00:01';ON CREATE CURRENT_TIMESTAMP" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"not null;default:'1970-01-01 00:00:01';ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt,omitempty"`
}

// BeforeCreate implement method BeforeCreate() from interface gorm.BeforeCreate on Note
func (n *Note) BeforeCreate(db *gorm.DB) (err error) {
	n.ID = uuid.New().String()
	return nil
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

// ValidateStruct create function to ValidateStruct[T any](payload T) with return of []*ErrorResponse
func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse

	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorResponse{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Param(),
			})
		}
	}

	return errors
}

// CreateNoteSchema struct of CreateNoteSchema
type CreateNoteSchema struct {
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Category  string    `json:"category,omitempty"`
	Published bool      `json:"published,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// UpdateNoteSchema create struct of UpdateNoteSchema
type UpdateNoteSchema struct {
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	Category  string    `json:"category,omitempty"`
	Published *bool     `json:"published,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
