package repository

import (
	"academic-system/internal/config"
	"academic-system/internal/models"
	"gorm.io/gorm"
)

type AssessmentRepository struct {
	db *gorm.DB
}

func NewAssessmentRepository() *AssessmentRepository {
	return &AssessmentRepository{db: config.DB}
}

func (r *AssessmentRepository) Create(assessment *models.Assessment) error {
	return r.db.Create(assessment).Error
}

func (r *AssessmentRepository) Update(assessment *models.Assessment) error {
	return r.db.Save(assessment).Error
}

func (r *AssessmentRepository) GetByID(id uint) (*models.Assessment, error) {
	var assessment models.Assessment
	err := r.db.Preload("Teacher").Preload("Class").Preload("Subject").
		   Preload("Students").First(&assessment, id).Error
	return &assessment, err
}

func (r *AssessmentRepository) List(filters map[string]interface{}) ([]models.Assessment, error) {
	var assessments []models.Assessment
	query := r.db.Preload("Teacher").Preload("Class").Preload("Subject")
	
	for key, value := range filters {
		query = query.Where(key, value)
	}
	
	err := query.Find(&assessments).Error
	return assessments, err
}