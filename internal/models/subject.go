package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Assessments []Assessment
}