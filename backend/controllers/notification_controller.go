package controllers

import (
	"net/http"
	"schoolapp/backend/config"
	"schoolapp/backend/models"
	"schoolapp/backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateNotificationInput struct {
	Message string `json:"message" binding:"required"`
}

func GetNotifications(c *gin.Context) {
	var notifications []models.Notification
	err := config.DB.Order("created_at desc").Limit(30).Find(&notifications).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal memuat notifikasi",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    notifications,
	})
}

func CreateNotification(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Anda tidak memiliki hak akses untuk menyiarkan notifikasi",
		})
		return
	}

	var input CreateNotificationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Pesan notifikasi wajib diisi",
		})
		return
	}

	notification := models.Notification{
		Message: input.Message,
		IsRead:  false,
	}

	err := config.DB.Create(&notification).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal menyimpan notifikasi",
		})
		return
	}

	if services.SocketServer != nil {
		services.SocketServer.BroadcastToNamespace("/", "notification", input.Message)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    notification,
	})
}

func ClearNotifications(c *gin.Context) {
	err := config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Notification{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal membersihkan notifikasi",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Semua notifikasi berhasil dibersihkan",
	})
}
