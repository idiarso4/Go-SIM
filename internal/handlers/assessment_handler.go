package handlers

import (
	"net/http"
	"strconv"
	"academic-system/internal/models"
	"academic-system/internal/repository"
	"github.com/gin-gonic/gin"
)

var assessmentRepo = repository.NewAssessmentRepository()

func CreateAssessment(c *gin.Context) {
	var assessment models.Assessment
	if err := c.ShouldBindJSON(&assessment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := assessmentRepo.Create(&assessment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, assessment)
}

func UpdateAssessment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var assessment models.Assessment
	if err := c.ShouldBindJSON(&assessment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assessment.ID = uint(id)
	if err := assessmentRepo.Update(&assessment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assessment)
}

func GetAssessment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	assessment, err := assessmentRepo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assessment not found"})
		return
	}

	c.JSON(http.StatusOK, assessment)
}

func ListAssessments(c *gin.Context) {
	filters := make(map[string]interface{})
	// Add filter parsing logic here
	
	assessments, err := assessmentRepo.List(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assessments)
}