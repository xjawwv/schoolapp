package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"schoolapp/backend/config"
	"schoolapp/backend/models"
	"schoolapp/backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetPopups(c *gin.Context) {
	var popups []models.Popup
	err := config.DB.Order("created_at desc").Limit(30).Find(&popups).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal memuat popup pengumuman",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    popups,
	})
}

func CreatePopup(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Anda tidak memiliki hak akses untuk menyiarkan popup",
		})
		return
	}

	title := c.PostForm("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Judul popup wajib diisi",
		})
		return
	}

	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "File foto popup wajib dikirim",
		})
		return
	}

	if file.Size > 3*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Ukuran file terlalu besar. Maksimum ukuran adalah 3MB",
		})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" && ext != ".gif" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Format file tidak didukung. Hanya diperbolehkan JPG, JPEG, PNG, WEBP, atau GIF",
		})
		return
	}

	contentType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Tipe file tidak valid. File harus berupa gambar",
		})
		return
	}

	uploadDir := filepath.Join(".", "uploads", "announcements")
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal membuat folder penyimpanan",
		})
		return
	}

	uniqueID := uuid.New().String()
	filename := uniqueID + ext
	filePath := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal menyimpan file foto popup",
		})
		return
	}

	photoPath := "/uploads/announcements/" + filename

	popup := models.Popup{
		Title: title,
		Photo: photoPath,
	}

	if err := config.DB.Create(&popup).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal menyimpan popup ke database",
		})
		return
	}

	if services.SocketServer != nil {
		services.SocketServer.BroadcastToNamespace("/", "popup", map[string]interface{}{
			"id":         popup.ID,
			"title":      popup.Title,
			"photo":      popup.Photo,
			"created_at": popup.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    popup,
	})
}

func ClearPopups(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Anda tidak memiliki hak akses untuk membersihkan popup",
		})
		return
	}

	err := config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Popup{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal membersihkan riwayat popup",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Semua popup berhasil dibersihkan",
	})
}
