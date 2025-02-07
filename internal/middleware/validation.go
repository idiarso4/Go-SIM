package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"academic-system/internal/models"
)

func ValidateClass() gin.HandlerFunc {
	return func(c *gin.Context) {
		var class models.Class
		if err := c.ShouldBindJSON(&class); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("validatedClass", class)
		c.Next()
	}
}

func ValidateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var student models.Student
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("validatedStudent", student)
		c.Next()
	}
}