package controllers

import (
	"net/http"
	"schoolapp/backend/config"
	"schoolapp/backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type PasswordInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type SettingsInput struct {
	SiteName string `json:"site_name" binding:"required"`
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
	var input SettingsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	var setting models.Setting
	err := config.DB.Where("key = ?", "site_name").First(&setting).Error
	if err != nil {
		setting = models.Setting{Key: "site_name", Value: input.SiteName}
		config.DB.Create(&setting)
	} else {
		setting.Value = input.SiteName
		config.DB.Save(&setting)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"site_name": input.SiteName,
		},
		"message": "Pengaturan berhasil disimpan",
	})
}

func ChangePassword(c *gin.Context) {
	var input PasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	userID := c.MustGet("userId").(string)

	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    nil,
			"message": "Pengguna tidak ditemukan",
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
