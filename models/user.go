package models

import (
	"gorm.io/gorm"
	"hacktiv8-final-project4/helpers"
)

const (
	AdminRole    string = "admin"
	CustomerRole string = "customer"
)

type User struct {
	GormModel
	FullName string `gorm:"column:full_name" json:"full_name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Role     string `gorm:"column:role" json:"role"`
	Balance  int    `gorm:"column:balance" json:"balance"`
}

func (user *User) VerifyPassword(password string) bool {
	return helpers.VerifyPassword(password, user.Password)
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Password = helpers.HashPassword(user.Password)

	// Assign default role if not set
	if user.Role == "" {
		user.Role = CustomerRole
	}
	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("password") {
		user.Password = helpers.HashPassword(user.Password)
	}
	return nil
}
