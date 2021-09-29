package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	FullName string
	Email    string
	Phone    string
	OIDCId   string `gorm:"column:oidc_id"`
	OIDCType string `gorm:"column:oidc_type"`
	Admin    bool
}
