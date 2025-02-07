package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	NIS         string `gorm:"unique;not null"`
	Name        string `not null`
	ClassID     uint
	Class       Class `gorm:"foreignKey:ClassID"`
	Assessments []Assessment `gorm:"many2many:student_assessments;"`
}