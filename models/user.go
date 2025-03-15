package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// ID       uint   `gorm:"primaryKey;auto_increment"`
	Email    string `json:"email" gorm:"unique"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"type:varchar(12)"`
}
