package controllers

import (
	"net/http"
	"strings"

	"schoolapp/backend/config"
	"schoolapp/backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type PasswordInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

func GetSettings(c *gin.Context) {
	var settings []models.Setting
	if err := config.DB.Find(&settings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal mengambil pengaturan",
		})
		return
	}

	result := make(map[string]string)
	for _, s := range settings {
		result[s.Key] = s.Value
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
		"message": "Pengaturan berhasil diambil",
	})
}

func UpdateSettings(c *gin.Context) {
	var input map[string]string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	for k, v := range input {
		var setting models.Setting
		err := config.DB.Where("key = ?", k).First(&setting).Error
		if err != nil {
			setting = models.Setting{Key: k, Value: v}
			config.DB.Create(&setting)
		} else {
			setting.Value = v
			config.DB.Save(&setting)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    input,
		"message": "Pengaturan berhasil disimpan",
	})
}

func ChangePassword(c *gin.Context) {
	var input PasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		message := "Input tidak valid"
		if strings.Contains(err.Error(), "min") {
			message = "Kata sandi baru harus minimal 6 karakter"
		} else if strings.Contains(err.Error(), "required") {
			message = "Kata sandi lama dan baru wajib diisi"
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": message,
		})
		return
	}

	userID := c.MustGet("userId").(string)

	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"data":    nil,
			"message": "Sesi tidak valid atau pengguna tidak ditemukan",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Password lama salah",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal memproses password baru",
		})
		return
	}

	user.Password = string(hashedPassword)
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Gagal menyimpan password baru",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    nil,
		"message": "Password berhasil diperbarui",
	})
}
