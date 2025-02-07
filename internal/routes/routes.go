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
		// Assessment routes
		assessment := protected.Group("/assessments")
		{
			assessment.POST("/", handlers.CreateAssessment)
			assessment.PUT("/:id", handlers.UpdateAssessment)
			assessment.GET("/:id", handlers.GetAssessment)
			assessment.GET("/", handlers.ListAssessments)
			assessment.POST("/:id/export", handlers.ExportAssessment)
		}

		// Class routes
		class := protected.Group("/classes")
		{
			class.GET("/", handlers.ListClasses)
			class.GET("/:id/students", handlers.GetClassStudents)
		}

		// Student routes
		student := protected.Group("/students")
		{
			student.GET("/", handlers.ListStudents)
			student.GET("/:id/assessments", handlers.GetStudentAssessments)
		}
	}
}
func SetupRoutes(r *gin.Engine) {
	// ... existing routes ...

	// Class routes
	classes := r.Group("/api/classes")
	{
		classes.POST("/", middleware.ValidateClass(), handlers.CreateClass)
		classes.GET("/:id", handlers.GetClass)
		classes.GET("/", handlers.ListClasses)
	}

	// Student routes
	students := r.Group("/api/students")
	{
		students.POST("/", middleware.ValidateStudent(), handlers.CreateStudent)
		students.GET("/:id", handlers.GetStudent)
		students.GET("/", handlers.ListStudents)
	}

	// Export route
	r.GET("/api/export/assessment/:id", handlers.ExportAssessment)
}
