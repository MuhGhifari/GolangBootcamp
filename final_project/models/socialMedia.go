package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Name           string `gorm:"not null" json:"name" form:"name" validate:"required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" validate:"required"`
	UserId         int    `json:"user_id" form:"user_id" validate:"numeric"`
	User           *User
}

func (sc *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errCreate := validate.Struct(sc)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}
