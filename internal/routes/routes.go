package routes

import (
    "academic-system/internal/handlers"
    "academic-system/internal/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    // Public routes
    public := r.Group("/api")
    {
        public.POST("/login", handlers.Login)
        public.POST("/register", handlers.Register)
    }

    // Protected routes
    protected := r.Group("/api")
    protected.Use(middleware.AuthMiddleware())
    {
        // Class routes
        class := protected.Group("/classes")
        {
            class.POST("/", middleware.ValidateClass(), handlers.CreateClass)
            class.PUT("/:id", middleware.ValidateClass(), handlers.UpdateClass)
            class.DELETE("/:id", handlers.DeleteClass)
            class.GET("/:id", handlers.GetClass)
            class.GET("/", handlers.ListClasses)
        }

        // Package routes
        package_ := protected.Group("/packages")
        {
            package_.POST("/", middleware.ValidatePackage(), handlers.CreatePackage)
            package_.PUT("/:id", middleware.ValidatePackage(), handlers.UpdatePackage)
            package_.DELETE("/:id", handlers.DeletePackage)
            package_.GET("/:id", handlers.GetPackage)
            package_.GET("/", handlers.ListPackages)
        }

        // Student routes
        student := protected.Group("/students")
        {
            student.POST("/", middleware.ValidateStudent(), handlers.CreateStudent)
            student.PUT("/:id", middleware.ValidateStudent(), handlers.UpdateStudent)
            student.DELETE("/:id", handlers.DeleteStudent)
            student.GET("/:id", handlers.GetStudent)
            student.GET("/", handlers.ListStudents)
        }

        // Assessment routes
        assessment := protected.Group("/assessments")
        {
            assessment.POST("/:id/export", handlers.ExportAssessment)
        }
    }
}