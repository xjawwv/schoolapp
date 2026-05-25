package routes

import (
	"net/http"

	"schoolapp/backend/controllers"
	"schoolapp/backend/middleware"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware())

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/register", controllers.Register)
		}

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/students", controllers.GetStudents)
			protected.GET("/students/:id", controllers.GetStudentByID)
			protected.POST("/students", controllers.CreateStudent)
			protected.PUT("/students/:id", controllers.UpdateStudent)
			protected.DELETE("/students/:id", controllers.DeleteStudent)

			protected.GET("/attendances", controllers.GetAttendances)
			protected.POST("/attendances", controllers.InputAttendance)
			protected.PUT("/attendances/:id", controllers.UpdateAttendance)

			protected.GET("/grades", controllers.GetGrades)
			protected.POST("/grades", controllers.InputGrade)
			protected.PUT("/grades/:id", controllers.UpdateGrade)

			protected.GET("/dashboard/stats", controllers.GetDashboardStats)

			protected.GET("/settings", controllers.GetSettings)
			protected.PUT("/settings", controllers.UpdateSettings)
			protected.PUT("/settings/password", controllers.ChangePassword)
		}
	}

	return r
}
