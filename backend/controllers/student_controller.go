package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"schoolapp/backend/config"
	"schoolapp/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type StudentInput struct {
	NISN         string `json:"nisn" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Class        string `json:"class" binding:"required"`
	Gender       string `json:"gender" binding:"required"`
	PlaceOfBirth string `json:"place_of_birth"`
	DateOfBirth  string `json:"date_of_birth"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
}

func generateStudentPassword(dateOfBirth string, nisn string) string {
	yearPart := "00"
	if len(dateOfBirth) >= 4 {
		year := strings.Split(dateOfBirth, "-")[0]
		if len(year) >= 4 {
			yearPart = year[len(year)-2:]
		}
	}
	nisn2 := "00"
	if len(nisn) >= 2 {
		nisn2 = nisn[len(nisn)-2:]
	}
	return fmt.Sprintf("SISWA%s%s", yearPart, nisn2)
}

func GetStudents(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	search := c.Query("search")

	var students []models.Student
	var total int64

	query := config.DB.Model(&models.Student{})

	if search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("name ILIKE ? OR nisn ILIKE ? OR class ILIKE ?", searchTerm, searchTerm, searchTerm)
	}

	class := c.Query("class")
	if class != "" {
		query = query.Where("class = ?", class)
	}

	query.Count(&total)

	offset := (page - 1) * limit
	err = query.Order("name ASC").Limit(limit).Offset(offset).Find(&students).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to fetch students",
		})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"students":    students,
			"total":       total,
			"page":        page,
			"limit":       limit,
			"total_pages": totalPages,
		},
		"message": "Students retrieved successfully",
	})
}

func GetStudentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Invalid student UUID format",
		})
		return
	}

	var student models.Student
	if err := config.DB.First(&student, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    nil,
			"message": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    student,
		"message": "Student retrieved successfully",
	})
}

func CreateStudent(c *gin.Context) {
	var input StudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	var existingStudent models.Student
	if err := config.DB.Where("nisn = ?", input.NISN).First(&existingStudent).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"data":    nil,
			"message": "NISN is already in use by another student",
		})
		return
	}

	student := models.Student{
		NISN:         input.NISN,
		Name:         input.Name,
		Class:        input.Class,
		Gender:       input.Gender,
		PlaceOfBirth: input.PlaceOfBirth,
		DateOfBirth:  input.DateOfBirth,
		Address:      input.Address,
		Phone:        input.Phone,
	}

	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to create student record",
		})
		return
	}

	defaultPwd := generateStudentPassword(input.DateOfBirth, input.NISN)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(defaultPwd), bcrypt.DefaultCost)
	role := "siswa"
	if student.Gender == "Perempuan" {
		role = "siswi"
	}
	user := models.User{
		Name:      student.Name,
		Email:     "student_" + student.NISN + "@sekolah.com",
		Password:  string(hashedPassword),
		Role:      role,
		StudentID: &student.ID,
	}
	if errDb := config.DB.Create(&user).Error; errDb != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal membuat akun pengguna siswa: " + errDb.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    student,
		"message": "Student created successfully",
	})
}

func UpdateStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Invalid student UUID format",
		})
		return
	}

	var student models.Student
	if err := config.DB.First(&student, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    nil,
			"message": "Student not found",
		})
		return
	}

	var input StudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	var existingStudent models.Student
	if err := config.DB.Where("nisn = ? AND id != ?", input.NISN, id).First(&existingStudent).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"data":    nil,
			"message": "NISN is already in use by another student",
		})
		return
	}

	student.NISN = input.NISN
	student.Name = input.Name
	student.Class = input.Class
	student.Gender = input.Gender
	student.PlaceOfBirth = input.PlaceOfBirth
	student.DateOfBirth = input.DateOfBirth
	student.Address = input.Address
	student.Phone = input.Phone

	if err := config.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to update student record",
		})
		return
	}

	var user models.User
	if err := config.DB.Where("student_id = ?", student.ID).First(&user).Error; err == nil {
		user.Name = student.Name
		user.Email = "student_" + student.NISN + "@sekolah.com"
		role := "siswa"
		if student.Gender == "Perempuan" {
			role = "siswi"
		}
		user.Role = role
		config.DB.Save(&user)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    student,
		"message": "Student updated successfully",
	})
}

func DeleteStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Invalid student UUID format",
		})
		return
	}

	var student models.Student
	if err := config.DB.First(&student, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    nil,
			"message": "Student not found",
		})
		return
	}

	tx := config.DB.Begin()

	if err := tx.Where("student_id = ?", id).Delete(&models.Attendance{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to clean up associated student attendance records",
		})
		return
	}

	if err := tx.Where("student_id = ?", id).Delete(&models.Grade{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to clean up associated student grade records",
		})
		return
	}

	if err := tx.Where("student_id = ?", id).Delete(&models.User{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to clean up associated student user account",
		})
		return
	}

	if err := tx.Delete(&student).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to delete student record",
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    nil,
		"message": "Student and all associated records deleted successfully",
	})
}
