package model

import (
	"gorm.io/gorm"
	"hacktiv8-final-project4/helper"
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
	return helper.VerifyPassword(password, user.Password)
}

func (user *User) GenerateJwtToken() string {
	return helper.GenerateToken(user.ID, user.Email, user.Role)
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Password = helper.HashPassword(user.Password)

	// Assign default role if not set
	if user.Role == "" {
		user.Role = CustomerRole
	}
	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("password") {
		user.Password = helper.HashPassword(user.Password)
	}
	return nil
}
