package controllers

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"schoolapp/backend/config"
	"schoolapp/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AttendanceInput struct {
	StudentID uuid.UUID `json:"student_id" binding:"required"`
	Date      string    `json:"date" binding:"required"`
	Status    string    `json:"status" binding:"required"`
	Note      string    `json:"note"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp string    `json:"timestamp"`
}

func GetAttendances(c *gin.Context) {
	var attendances []models.Attendance

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
				"data":    []models.Attendance{},
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

	dateStr := c.Query("date")
	if dateStr != "" {
		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err == nil {
			query = query.Where("date = ?", parsedDate)
		}
	}

	joinedStudent := false

	class := c.Query("class")
	if class != "" {
		query = query.Joins("JOIN students ON students.id = attendances.student_id")
		joinedStudent = true
		query = query.Where("students.class = ?", class)
	}

	search := c.Query("search")
	if search != "" {
		if !joinedStudent {
			query = query.Joins("JOIN students ON students.id = attendances.student_id")
		}
		query = query.Where("students.name ILIKE ? OR students.nisn ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	err := query.Order("date DESC").Find(&attendances).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to fetch attendance records",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    attendances,
		"message": "Attendance records retrieved successfully",
	})
}

func InputAttendance(c *gin.Context) {
	var input AttendanceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	if input.Timestamp == "" {
		input.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	}

	parsedDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Date must be in YYYY-MM-DD format",
		})
		return
	}

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
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"data":    nil,
				"message": "Akun siswa belum terhubung dengan profil siswa",
			})
			return
		}
		if input.StudentID != *user.StudentID {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"data":    nil,
				"message": "Akses ditolak: Anda hanya dapat mengisi kehadiran diri Anda sendiri",
			})
			return
		}
		if parsedDate.Format("2006-01-02") != time.Now().Format("2006-01-02") {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"data":    nil,
				"message": "Akses ditolak: Anda hanya diperbolehkan melakukan absensi mandiri untuk hari ini",
			})
			return
		}

		var startSetting models.Setting
		startTimeStr := "07:00"
		if err := config.DB.Where("key = ?", "attendance_start_time").First(&startSetting).Error; err == nil {
			startTimeStr = startSetting.Value
		}

		var endSetting models.Setting
		endTimeStr := "17:00"
		if err := config.DB.Where("key = ?", "attendance_end_time").First(&endSetting).Error; err == nil {
			endTimeStr = endSetting.Value
		}

		nowTimeStr := time.Now().Format("15:04")
		if nowTimeStr < startTimeStr || nowTimeStr > endTimeStr {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"data":    nil,
				"message": "Absensi ditolak: Waktu absensi mandiri hanya diizinkan pukul " + startTimeStr + " s.d. " + endTimeStr,
			})
			return
		}

		if input.Status == "hadir" {
			if input.Latitude == 0 || input.Longitude == 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"data":    nil,
					"message": "Lokasi GPS diperlukan untuk status kehadiran Hadir",
				})
				return
			}

			var latSetting models.Setting
			schoolLat := -6.1822
			if err := config.DB.Where("key = ?", "school_latitude").First(&latSetting).Error; err == nil {
				fmt.Sscanf(latSetting.Value, "%f", &schoolLat)
			}

			var lonSetting models.Setting
			schoolLon := 106.2736
			if err := config.DB.Where("key = ?", "school_longitude").First(&lonSetting).Error; err == nil {
				fmt.Sscanf(lonSetting.Value, "%f", &schoolLon)
			}

			var distSetting models.Setting
			maxDistance := 250.0
			if err := config.DB.Where("key = ?", "school_max_distance").First(&distSetting).Error; err == nil {
				fmt.Sscanf(distSetting.Value, "%f", &maxDistance)
			}

			dist := calculateDistance(input.Latitude, input.Longitude, schoolLat, schoolLon)
			if dist > maxDistance {
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"data":    nil,
					"message": "Anda berada di luar jangkauan area sekolah (jarak Anda: " + fmt.Sprintf("%.1f", dist) + " meter). Absensi ditolak",
				})
				return
			}
		}
	}

	validStatuses := map[string]bool{"hadir": true, "sakit": true, "izin": true, "alpha": true}
	if !validStatuses[input.Status] {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Status must be either 'hadir', 'sakit', 'izin', or 'alpha'",
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

	var attendance models.Attendance
	err = config.DB.Where("student_id = ? AND date = ?", input.StudentID, parsedDate).First(&attendance).Error
	if err == nil {
		attendance.Status = input.Status
		attendance.Note = input.Note
		attendance.Latitude = input.Latitude
		attendance.Longitude = input.Longitude
		attendance.Timestamp = input.Timestamp
		if err := config.DB.Save(&attendance).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"data":    nil,
				"message": "Failed to update attendance record",
			})
			return
		}

		config.DB.Preload("Student").First(&attendance, attendance.ID)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    attendance,
			"message": "Attendance record updated successfully",
		})
		return
	}

	attendance = models.Attendance{
		StudentID: input.StudentID,
		Date:      parsedDate,
		Status:    input.Status,
		Note:      input.Note,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		Timestamp: input.Timestamp,
	}

	if err := config.DB.Create(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to create attendance record",
		})
		return
	}

	config.DB.Preload("Student").First(&attendance, attendance.ID)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    attendance,
		"message": "Attendance record entered successfully",
	})
}

func UpdateAttendance(c *gin.Context) {
	role, roleExists := c.Get("role")
	if roleExists && (role == "siswa" || role == "siswi") {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"data":    nil,
			"message": "Akses ditolak: Siswa tidak diperbolehkan mengubah catatan kehadiran",
		})
		return
	}

	id := c.Param("id")

	var attendance models.Attendance
	if err := config.DB.First(&attendance, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    nil,
			"message": "Attendance record not found",
		})
		return
	}

	var input struct {
		Status string `json:"status" binding:"required"`
		Note   string `json:"note"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	validStatuses := map[string]bool{"hadir": true, "sakit": true, "izin": true, "alpha": true}
	if !validStatuses[input.Status] {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "Status must be either 'hadir', 'sakit', 'izin', or 'alpha'",
		})
		return
	}

	attendance.Status = input.Status
	attendance.Note = input.Note

	if err := config.DB.Save(&attendance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "Failed to update attendance record",
		})
		return
	}

	config.DB.Preload("Student").First(&attendance, attendance.ID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    attendance,
		"message": "Attendance record updated successfully",
	})
}

func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const EarthRadiusKm = 6371.0
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180.0)*math.Cos(lat2*math.Pi/180.0)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return EarthRadiusKm * c * 1000.0
}
