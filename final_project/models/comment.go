package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Message   string     `gorm:"not null" json:"message" form:"message" validate:"required"`
	UserId    int        `json:"user_id" form:"user_id" validate:"numeric"`
	PhotoId   int        `json:"photo_id" form:"photo_id" validate:"numeric"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	User      *User
	Photo     *Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errCreate := validate.Struct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errCreate := validate.Struct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}
