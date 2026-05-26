package controllers

import (
	"net/http"

	"schoolapp/backend/config"
	"schoolapp/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GradeInput struct {
	StudentID    uuid.UUID `json:"student_id" binding:"required"`
	Subject      string    `json:"subject" binding:"required"`
	Score        float64   `json:"score" binding:"required"`
	Semester     int       `json:"semester" binding:"required,min=1,max=2"`
	AcademicYear string    `json:"academic_year" binding:"required"`
}

func GetGrades(c *gin.Context) {
	var grades []models.Grade

	query := config.DB.Preload("Student")

	role, roleExists := c.Get("role")
	if roleExists && (role == "siswa" || role == "siswi") {
		userIdStr, userExists := c.Get("userId")
		if !userExists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"data":    nil,
				"message": "Sesi tidak valid",
			})
			return
		}
		var user models.User
		if err := config.DB.First(&user, "id = ?", userIdStr).Error; err != nil || user.StudentID == nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"data":    []models.Grade{},
				"message": "Akun siswa belum terhubung dengan profil siswa",
			})
			return
		}
		query = query.Where("student_id = ?", *user.StudentID)
	}

	studentID := c.Query("student_id")
	if studentID != "" {
		parsedUUID, err := uuid.Parse(studentID)
		if err == nil {
			query = query.Where("student_id = ?", parsedUUID)
		}
	}

	err := query.Order("created_at DESC").Find(&grades).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to fetch grade records",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    grades,
		"message": "Grade records retrieved successfully",
	})
}

func InputGrade(c *gin.Context) {
	role, roleExists := c.Get("role")
	if !roleExists || role != "guru" {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"data":    nil,
			"message": "Akses ditolak: Hanya Guru yang diperbolehkan menginput nilai baru",
		})
		return
	}

	var input GradeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	if input.Score < 0 || input.Score > 100 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Score must be between 0 and 100",
		})
		return
	}

	var student models.Student
	if err := config.DB.First(&student, "id = ?", input.StudentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    nil,
			"message": "Student not found",
		})
		return
	}

	var grade models.Grade
	err := config.DB.Where("student_id = ? AND subject = ? AND semester = ? AND academic_year = ?", input.StudentID, input.Subject, input.Semester, input.AcademicYear).First(&grade).Error
	if err == nil {
		grade.Score = input.Score
		if err := config.DB.Save(&grade).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"data":    nil,
				"message": "Failed to update grade record",
			})
			return
		}

		config.DB.Preload("Student").First(&grade, grade.ID)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    grade,
			"message": "Grade record updated successfully",
		})
		return
	}

	grade = models.Grade{
		StudentID:    input.StudentID,
		Subject:      input.Subject,
		Score:        input.Score,
		Semester:     input.Semester,
		AcademicYear: input.AcademicYear,
	}

	if err := config.DB.Create(&grade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to create grade record",
		})
		return
	}

	config.DB.Preload("Student").First(&grade, grade.ID)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    grade,
		"message": "Grade record entered successfully",
	})
}

func UpdateGrade(c *gin.Context) {
	role, roleExists := c.Get("role")
	if !roleExists || (role != "guru" && role != "admin") {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"data":    nil,
			"message": "Akses ditolak: Hanya Guru dan Admin yang diperbolehkan memperbaiki nilai",
		})
		return
	}

	id := c.Param("id")

	var grade models.Grade
	if err := config.DB.First(&grade, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    nil,
			"message": "Grade record not found",
		})
		return
	}

	var input struct {
		Score        float64 `json:"score" binding:"required"`
		Subject      string  `json:"subject" binding:"required"`
		Semester     int     `json:"semester" binding:"required,min=1,max=2"`
		AcademicYear string  `json:"academic_year" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	if input.Score < 0 || input.Score > 100 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Score must be between 0 and 100",
		})
		return
	}

	grade.Score = input.Score
	grade.Subject = input.Subject
	grade.Semester = input.Semester
	grade.AcademicYear = input.AcademicYear

	if err := config.DB.Save(&grade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to update grade record",
		})
		return
	}

	config.DB.Preload("Student").First(&grade, grade.ID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    grade,
		"message": "Grade record updated successfully",
	})
}
