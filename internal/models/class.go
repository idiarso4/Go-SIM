package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name      string    `gorm:"unique;not null"`
	Students  []Student
}