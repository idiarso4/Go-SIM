package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "academic-system/internal/models"
)

var validate = validator.New()

func ValidateClass() gin.HandlerFunc {
    return func(c *gin.Context) {
        var class models.Class
        if err := c.ShouldBindJSON(&class); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        if err := validate.Struct(&class); err != nil {
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

        if err := validate.Struct(&student); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        c.Set("validatedStudent", student)
        c.Next()
    }
}

func ValidateAssessment() gin.HandlerFunc {
    return func(c *gin.Context) {
        var assessment models.Assessment
        if err := c.ShouldBindJSON(&assessment); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        if err := validate.Struct(&assessment); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        c.Set("validatedAssessment", assessment)
        c.Next()
    }
}