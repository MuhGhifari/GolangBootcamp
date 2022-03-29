package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Title     string     `gorm:"not null" json:"title" form:"title" validate:"required"`
	Caption   string     `json:"caption" form:"caption"`
	PhotoUrl  string     `gorm:"not null" json:"photo_url" form:"photo_url" validate:"required"`
	UserId    int        `json:"user_id" form:"user_id" validate:"numeric"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Comment   []Comment  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	User      *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errCreate := validate.Struct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errCreate := validate.Struct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}
