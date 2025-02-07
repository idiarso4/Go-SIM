package models

import "gorm.io/gorm"

// models/student.go
type Student struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    NIS       string    `json:"nis" validate:"required,min=5,max=20"`
    Name      string    `json:"name" validate:"required,min=3,max=100"`
    ClassID   uint      `json:"class_id" validate:"required"`
    Class     Class     `json:"class"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`