package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"schoolapp/backend/config"
	"schoolapp/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Name      string     `json:"name" binding:"required"`
	Email     string     `json:"email" binding:"required,email"`
	Password  string     `json:"password" binding:"required,min=6"`
	Role      string     `json:"role" binding:"required"`
	StudentID *uuid.UUID `json:"student_id"`
}

type UpdateUserInput struct {
	Name      string     `json:"name" binding:"required"`
	Email     string     `json:"email" binding:"required,email"`
	Role      string     `json:"role" binding:"required"`
	StudentID *uuid.UUID `json:"student_id"`
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

	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"data":    nil,
			"message": "Email sudah terdaftar",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal memproses password",
		})
		return
	}

	user := models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  string(hashedPassword),
		Role:      input.Role,
		StudentID: input.StudentID,
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
