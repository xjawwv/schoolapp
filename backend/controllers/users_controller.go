package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"schoolapp/backend/config"
	"schoolapp/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Name         string     `json:"name" binding:"required"`
	Email        string     `json:"email" binding:"required,email"`
	Password     string     `json:"password"`
	Role         string     `json:"role" binding:"required"`
	StudentID    *uuid.UUID `json:"student_id"`
	NIP          string     `json:"nip"`
	PlaceOfBirth string     `json:"place_of_birth"`
	DateOfBirth  string     `json:"date_of_birth"`
	Gender       string     `json:"gender"`
	Address      string     `json:"address"`
	WhatsApp     string     `json:"whatsapp"`
}

type UpdateUserInput struct {
	Name         string     `json:"name" binding:"required"`
	Email        string     `json:"email" binding:"required,email"`
	Role         string     `json:"role" binding:"required"`
	StudentID    *uuid.UUID `json:"student_id"`
	NIP          string     `json:"nip"`
	PlaceOfBirth string     `json:"place_of_birth"`
	DateOfBirth  string     `json:"date_of_birth"`
	Gender       string     `json:"gender"`
	Address      string     `json:"address"`
	WhatsApp     string     `json:"whatsapp"`
}

func generateTeacherPassword(dateOfBirth string, nip string) string {
	yearPart := "00"
	if len(dateOfBirth) >= 4 {
		year := strings.Split(dateOfBirth, "-")[0]
		if len(year) >= 4 {
			yearPart = year[len(year)-2:]
		}
	}
	nip2 := "00"
	if len(nip) >= 2 {
		nip2 = nip[len(nip)-2:]
	}
	return fmt.Sprintf("GURU%s%s", yearPart, nip2)
}

func GetUsers(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"data":    nil,
			"message": "Akses ditolak",
		})
		return
	}

	var users []models.User
	if err := config.DB.Preload("Student").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal mengambil data pengguna",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
		"message": "Data pengguna berhasil diambil",
	})
}

func CreateUser(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"data":    nil,
			"message": "Akses ditolak",
		})
		return
	}

	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	if input.Role != "admin" && input.Role != "guru" && input.Role != "siswa" && input.Role != "siswi" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Role tidak valid",
		})
		return
	}

	if input.Role == "guru" && input.NIP == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "NIP wajib diisi untuk peran Guru",
		})
		return
	}

	if (input.Role == "siswa" || input.Role == "siswi") && input.StudentID == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Biodata Siswa (NISN) wajib dihubungkan untuk peran Siswa/Siswi",
		})
		return
	}

	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"data":    nil,
			"message": "Email sudah terdaftar",
		})
		return
	}

	password := input.Password
	if input.Role == "guru" {
		password = generateTeacherPassword(input.DateOfBirth, input.NIP)
	}
	if password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Password wajib diisi",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal memproses password",
		})
		return
	}

	user := models.User{
		Name:         input.Name,
		Email:        input.Email,
		Password:     string(hashedPassword),
		Role:         input.Role,
		StudentID:    input.StudentID,
		NIP:          input.NIP,
		PlaceOfBirth: input.PlaceOfBirth,
		DateOfBirth:  input.DateOfBirth,
		Gender:       input.Gender,
		Address:      input.Address,
		WhatsApp:     input.WhatsApp,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal membuat pengguna baru",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    user,
		"message": "Pengguna baru berhasil dibuat",
	})
}

func UpdateUser(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"data":    nil,
			"message": "Akses ditolak",
		})
		return
	}

	idStr := c.Param("id")
	userId, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "ID pengguna tidak valid",
		})
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	if input.Role != "admin" && input.Role != "guru" && input.Role != "siswa" && input.Role != "siswi" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Role tidak valid",
		})
		return
	}

	if input.Role == "guru" && input.NIP == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "NIP wajib diisi untuk peran Guru",
		})
		return
	}

	if (input.Role == "siswa" || input.Role == "siswi") && input.StudentID == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Biodata Siswa (NISN) wajib dihubungkan untuk peran Siswa/Siswi",
		})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    nil,
			"message": "Pengguna tidak ditemukan",
		})
		return
	}

	var duplicateUser models.User
	if err := config.DB.Where("email = ? AND id != ?", input.Email, user.ID).First(&duplicateUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"data":    nil,
			"message": "Email sudah digunakan oleh akun lain",
		})
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Role = input.Role
	user.StudentID = input.StudentID
	user.NIP = input.NIP
	user.PlaceOfBirth = input.PlaceOfBirth
	user.DateOfBirth = input.DateOfBirth
	user.Gender = input.Gender
	user.Address = input.Address
	user.WhatsApp = input.WhatsApp

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal memperbarui data pengguna",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
		"message": "Data pengguna berhasil diperbarui",
	})
}

func DeleteUser(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"data":    nil,
			"message": "Akses ditolak",
		})
		return
	}

	idStr := c.Param("id")
	userId, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "ID pengguna tidak valid",
		})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    nil,
			"message": "Pengguna tidak ditemukan",
		})
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal menghapus pengguna",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    nil,
		"message": "Pengguna berhasil dihapus",
	})
}

func UploadAvatar(c *gin.Context) {
	currentUserIdStr, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"data":    nil,
			"message": "Sesi tidak valid",
		})
		return
	}

	var user models.User
	if err := config.DB.First(&user, "id = ?", currentUserIdStr).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    nil,
			"message": "Pengguna tidak ditemukan",
		})
		return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "File foto profil wajib dikirim",
		})
		return
	}

	uploadDir := filepath.Join(".", "uploads", "avatars")
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal membuat folder penyimpanan",
		})
		return
	}

	uniqueID := uuid.New().String()
	filename := uniqueID + filepath.Ext(file.Filename)
	filePath := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal menyimpan file",
		})
		return
	}

	avatarPath := "/uploads/avatars/" + filename
	user.Avatar = avatarPath

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal memperbarui foto profil",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"avatar": avatarPath,
		},
		"message": "Foto profil berhasil diperbarui",
	})
}

func GetTeachers(c *gin.Context) {
	var teachers []models.User
	if err := config.DB.Where("role = ?", "guru").Find(&teachers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal mengambil data guru",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    teachers,
		"message": "Daftar guru berhasil diambil",
	})
}

func GetTeacherByID(c *gin.Context) {
	idStr := c.Param("id")
	teacherId, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Format ID guru tidak valid",
		})
		return
	}

	var teacher models.User
	if err := config.DB.Where("id = ? AND role = ?", teacherId, "guru").First(&teacher).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    nil,
			"message": "Data guru tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    teacher,
		"message": "Data guru berhasil diambil",
	})
}

