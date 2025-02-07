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

func UpdateClass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var class models.Class
	if v, exists := c.Get("validatedClass"); exists {
		class = v.(models.Class)
	}

	if err := classRepo.Update(uint(id), &class); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, class)
}

func DeleteClass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := classRepo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}

func GetClass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

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