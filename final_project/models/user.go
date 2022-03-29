package models

import (
	"github.com/MuhGhifari/GolangBootcamp/final-project/helpers"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	// ID          uint          `json:"id" gorm:"primaryKey"`
	// Username    string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required"`
	// Email       string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid email format"`
	// Password    string        `gorm:"not null" json:"password" form:"password" valid:"required~Password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	// Age         int           `gorm:"not null" json:"age" form:"age" valid:"required~Age is required, type(int)"`
	// CreatedAt   *time.Time    `json:"created_at,omitempty"`
	// UpdatedAt   *time.Time    `json:"updated_at,omitempty"`
	// Photos      []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	// Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	// SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socialmedias"`
	GormModel
	Username    string        `gorm:"not null;unique" json:"username" form:"username" validate:"required"`
	Email       string        `gorm:"not null;uniqueIndex" json:"email" form:"email" validate:"required"`
	Password    string        `gorm:"not null" json:"password" form:"password" validate:"required,min=6"`
	Age         int           `gorm:"not null" json:"age" form:"age" validate:"required,numeric,min=9"`
	Photos      []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socialmedias"`
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
