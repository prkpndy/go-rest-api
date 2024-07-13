package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primary_key"`
	Email    string `json:"email" gorm:"unique;not null;"`
	Age      int    `json:"age" gorm:"not null;"`
	Username string `json:"username" gorm:"unique;not null;"`
	Password string `json:"password" gorm:"not null;"`
}
