package handlers

import (
	"net/http"
	"strconv"
	"academic-system/internal/models"
	"academic-system/internal/repository"
	"github.com/gin-gonic/gin"
)

var classRepo = repository.NewClassRepository()

func CreateClass(c *gin.Context) {
	var class models.Class
	if v, exists := c.Get("validatedClass"); exists {
		class = v.(models.Class)
	}

	if err := classRepo.Create(&class); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, class)
}

func GetClass(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	class, err := classRepo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}
	c.JSON(http.StatusOK, class)
}

func ListClasses(c *gin.Context) {
	classes, err := classRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, classes)
}