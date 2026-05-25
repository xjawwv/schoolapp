package controllers

import (
	"net/http"
	"time"

	"schoolapp/backend/config"
	"schoolapp/backend/models"

	"github.com/gin-gonic/gin"
)

func GetDashboardStats(c *gin.Context) {
	var totalStudents int64
	if err := config.DB.Model(&models.Student{}).Count(&totalStudents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to count total students",
		})
		return
	}

	now := time.Now()
	firstDayOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)

	var totalAttendances int64
	var presentAttendances int64

	config.DB.Model(&models.Attendance{}).Where("date >= ? AND date <= ?", firstDayOfMonth, lastDayOfMonth).Count(&totalAttendances)
	config.DB.Model(&models.Attendance{}).Where("date >= ? AND date <= ? AND status = ?", firstDayOfMonth, lastDayOfMonth, "hadir").Count(&presentAttendances)

	attendanceRate := 0.0
	if totalAttendances > 0 {
		attendanceRate = (float64(presentAttendances) / float64(totalAttendances)) * 100.0
	}

	var averageScore float64
	config.DB.Model(&models.Grade{}).Select("COALESCE(AVG(score), 0)").Row().Scan(&averageScore)

	var recentStudents []models.Student
	config.DB.Order("created_at DESC").Limit(5).Find(&recentStudents)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"total_students":   totalStudents,
			"attendance_rate":  attendanceRate,
			"average_score":    averageScore,
			"recent_students": recentStudents,
		},
		"message": "Dashboard statistics retrieved successfully",
	})
}
