package handlers

import (
	"net/http"
	"strconv"
	"academic-system/internal/models"
	"academic-system/internal/repository"
	"github.com/gin-gonic/gin"
)

var studentRepo = repository.NewStudentRepository()

func CreateStudent(c *gin.Context) {
	var student models.Student
	if v, exists := c.Get("validatedStudent"); exists {
		student = v.(models.Student)
	}

	if err := studentRepo.Create(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, student)
}

func GetStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	student, err := studentRepo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func ListStudents(c *gin.Context) {
	students, err := studentRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}