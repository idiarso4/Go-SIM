package models

import (
	"time"
	"gorm.io/gorm"
)

type Assessment struct {
	gorm.Model
	Date        time.Time
	TeacherID   uint
	Teacher     User      `gorm:"foreignKey:TeacherID"`
	ClassID     uint
	Class       Class     `gorm:"foreignKey:ClassID"`
	SubjectID   uint
	Subject     Subject   `gorm:"foreignKey:SubjectID"`
	Type        string    `gorm:"type:enum('Summative','Non-Summative')"`
	Category    string    `gorm:"type:enum('Theory','Practical')"`
	Notes       string
	Students    []Student `gorm:"many2many:student_assessments;"`
}

type StudentAssessment struct {
	gorm.Model
	StudentID    uint
	AssessmentID uint
	Score        float64 `gorm:"type:decimal(5,2);check:score >= 0 AND score <= 100"`
	Remarks      string
}