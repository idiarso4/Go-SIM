package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"unique;not null"`
	Password    string `not null`
	Role        string `gorm:"default:'teacher'"`
	FullName    string
	Email       string `gorm:"unique"`
	LastLogin   *time.Time
}