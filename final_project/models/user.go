package models

import (
	"github.com/MuhGhifari/GolangBootcamp/final-project/helpers"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string        `gorm:"not null;unique" json:"username" form:"username" validate:"required"`
	Email       string        `gorm:"not null;uniqueIndex" json:"email" form:"email" validate:"required,email"`
	Password    string        `gorm:"not null" json:"password" form:"password" validate:"required,min=6"`
	Age         int           `gorm:"not null" json:"age" form:"age" validate:"required,numeric,min=9"`
	Photos      []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socialmedias"`
}

type userUpdate struct {
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" from:"username" validate:"required,email"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	errCreate := validate.Struct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	data := userUpdate{
		Username: u.Username,
		Email: u.Email,
	}

	validate := validator.New()
	errUpdate := validate.Struct(data); if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
