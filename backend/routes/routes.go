package routes

import (
	"net/http"

	"schoolapp/backend/controllers"
	"schoolapp/backend/middleware"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
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

func SetupRouter(socketServer *socketio.Server) *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware())
	r.Any("/socket.io/*any", func(c *gin.Context) {
		socketServer.ServeHTTP(c.Writer, c.Request)
	})
	r.Static("/uploads", "./uploads")

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/register", controllers.Register)
		}

		api.GET("/settings", controllers.GetSettings)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/notifications", controllers.GetNotifications)
			protected.POST("/notifications", controllers.CreateNotification)
			protected.DELETE("/notifications", controllers.ClearNotifications)
			protected.GET("/popups", controllers.GetPopups)
			protected.POST("/popups", controllers.CreatePopup)
			protected.DELETE("/popups", controllers.ClearPopups)
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

			protected.GET("/users", controllers.GetUsers)
			protected.POST("/users", controllers.CreateUser)
			protected.PUT("/users/:id", controllers.UpdateUser)
			protected.DELETE("/users/:id", controllers.DeleteUser)
			protected.POST("/profile/avatar", controllers.UploadAvatar)
			protected.DELETE("/profile/avatar", controllers.DeleteAvatar)
			protected.GET("/teachers", controllers.GetTeachers)
			protected.GET("/teachers/:id", controllers.GetTeacherByID)

			protected.GET("/dashboard/stats", controllers.GetDashboardStats)

			protected.PUT("/settings", controllers.UpdateSettings)
			protected.PUT("/settings/password", controllers.ChangePassword)
		}
	}

	return r
}
