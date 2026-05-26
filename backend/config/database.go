package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"schoolapp/backend/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database

	DB.Exec(`
		DO $$ 
		BEGIN 
			IF EXISTS (
				SELECT 1 
				FROM information_schema.columns 
				WHERE table_name='students' AND column_name='nis'
			) THEN
				ALTER TABLE students RENAME COLUMN nis TO nisn;
			END IF;
		END $$;
	`)

	DB.Exec(`ALTER TABLE users ADD COLUMN IF NOT EXISTS nip varchar(50);`)
	DB.Exec(`ALTER TABLE users ADD COLUMN IF NOT EXISTS student_id uuid;`)
	DB.Exec(`ALTER TABLE users ADD COLUMN IF NOT EXISTS avatar varchar(255);`)

	err = DB.AutoMigrate(&models.User{}, &models.Student{}, &models.Attendance{}, &models.Grade{}, &models.Setting{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	DB.Exec("DROP INDEX IF EXISTS idx_users_nip;")
	DB.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_users_nip ON users(nip) WHERE nip IS NOT NULL AND nip <> '';")

	seedData()
}

func seedData() {
	var settingCount int64
	DB.Model(&models.Setting{}).Where("key = ?", "site_name").Count(&settingCount)
	if settingCount == 0 {
		siteName := models.Setting{
			Key:   "site_name",
			Value: "SMA N 1 METRO",
		}
		DB.Create(&siteName)
	}

	var admin models.User
	if err := DB.Unscoped().Where("email = ?", "admin@sekolah.com").First(&admin).Error; err != nil {
		hashedAdminPassword, errAdmin := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		if errAdmin == nil {
			admin := models.User{
				Name:     "Administrator",
				Email:    "admin@sekolah.com",
				Password: string(hashedAdminPassword),
				Role:     "admin",
			}
			DB.Create(&admin)
		}
	} else if !admin.DeletedAt.Time.IsZero() {
		admin.DeletedAt = gorm.DeletedAt{}
		DB.Unscoped().Save(&admin)
	}

	var guru models.User
	if err := DB.Unscoped().Where("email = ?", "guru@sekolah.com").First(&guru).Error; err == nil {
		var updated bool
		if guru.NIP != "12345" {
			guru.NIP = "12345"
			updated = true
		}
		if !guru.DeletedAt.Time.IsZero() {
			guru.DeletedAt = gorm.DeletedAt{}
			updated = true
		}
		if updated {
			DB.Unscoped().Save(&guru)
		}
	} else {
		hashedGuruPassword, errGuru := bcrypt.GenerateFromPassword([]byte("gurusmk"), bcrypt.DefaultCost)
		if errGuru == nil {
			guru := models.User{
				Name:     "Guru SMK",
				Email:    "guru@sekolah.com",
				NIP:      "12345",
				Password: string(hashedGuruPassword),
				Role:     "guru",
			}
			DB.Create(&guru)
		}
	}

	var studentCount int64
	DB.Model(&models.Student{}).Count(&studentCount)
	if studentCount == 0 {
		students := []models.Student{
			{NISN: "10001", Name: "Ahmad Fauzi", Class: "X-A", Gender: "Laki-laki", Address: "Jl. Merdeka No. 12", Phone: "081234567890"},
			{NISN: "10002", Name: "Budi Santoso", Class: "X-A", Gender: "Laki-laki", Address: "Jl. Mawar No. 4", Phone: "081234567891"},
			{NISN: "10003", Name: "Citra Lestari", Class: "X-A", Gender: "Perempuan", Address: "Jl. Melati No. 8", Phone: "081234567892"},
			{NISN: "10004", Name: "Dwi Cahyo", Class: "XI-B", Gender: "Laki-laki", Address: "Jl. Dahlia No. 15", Phone: "081234567893"},
			{NISN: "10005", Name: "Eka Putri", Class: "XI-B", Gender: "Perempuan", Address: "Jl. Anggrek No. 20", Phone: "081234567894"},
			{NISN: "10006", Name: "Fitriani", Class: "XI-B", Gender: "Perempuan", Address: "Jl. Tulip No. 3", Phone: "081234567895"},
			{NISN: "10007", Name: "Gede Wahyu", Class: "XII-C", Gender: "Laki-laki", Address: "Jl. Cempaka No. 7", Phone: "081234567896"},
			{NISN: "10008", Name: "Haryono", Class: "XII-C", Gender: "Laki-laki", Address: "Jl. Kamboja No. 9", Phone: "081234567897"},
			{NISN: "10009", Name: "Indah Permata", Class: "XII-C", Gender: "Perempuan", Address: "Jl. Kenanga No. 2", Phone: "081234567898"},
			{NISN: "10010", Name: "Joko Susilo", Class: "XII-C", Gender: "Laki-laki", Address: "Jl. Teratai No. 11", Phone: "081234567899"},
		}

		for i := range students {
			DB.Create(&students[i])
			hashedPassword, errPass := bcrypt.GenerateFromPassword([]byte("1"), bcrypt.DefaultCost)
			if errPass == nil {
				role := "siswa"
				if students[i].Gender == "Perempuan" {
					role = "siswi"
				}
				user := models.User{
					Name:      students[i].Name,
					Email:     "student_" + students[i].NISN + "@sekolah.com",
					Password:  string(hashedPassword),
					Role:      role,
					StudentID: &students[i].ID,
				}
				DB.Create(&user)
			}
		}

		var createdStudents []models.Student
		DB.Find(&createdStudents)

		if len(createdStudents) >= 5 {
			today := time.Now()
			yesterday := today.AddDate(0, 0, -1)

			attendances := []models.Attendance{
				{StudentID: createdStudents[0].ID, Date: today, Status: "hadir", Note: ""},
				{StudentID: createdStudents[1].ID, Date: today, Status: "hadir", Note: ""},
				{StudentID: createdStudents[2].ID, Date: today, Status: "sakit", Note: "Demam tinggi"},
				{StudentID: createdStudents[3].ID, Date: today, Status: "izin", Note: "Acara keluarga"},
				{StudentID: createdStudents[4].ID, Date: today, Status: "hadir", Note: ""},
				{StudentID: createdStudents[0].ID, Date: yesterday, Status: "hadir", Note: ""},
				{StudentID: createdStudents[1].ID, Date: yesterday, Status: "hadir", Note: ""},
				{StudentID: createdStudents[2].ID, Date: yesterday, Status: "hadir", Note: ""},
				{StudentID: createdStudents[3].ID, Date: yesterday, Status: "hadir", Note: ""},
				{StudentID: createdStudents[4].ID, Date: yesterday, Status: "alpha", Note: "Tanpa keterangan"},
			}
			for i := range attendances {
				DB.Create(&attendances[i])
			}

			grades := []models.Grade{
				{StudentID: createdStudents[0].ID, Subject: "Matematika", Score: 85.5, Semester: 1, AcademicYear: "2025/2026"},
				{StudentID: createdStudents[0].ID, Subject: "Fisika", Score: 78.0, Semester: 1, AcademicYear: "2025/2026"},
				{StudentID: createdStudents[1].ID, Subject: "Matematika", Score: 90.0, Semester: 1, AcademicYear: "2025/2026"},
				{StudentID: createdStudents[1].ID, Subject: "Bahasa Indonesia", Score: 88.5, Semester: 1, AcademicYear: "2025/2026"},
				{StudentID: createdStudents[2].ID, Subject: "Matematika", Score: 65.0, Semester: 1, AcademicYear: "2025/2026"},
				{StudentID: createdStudents[2].ID, Subject: "Kimia", Score: 72.5, Semester: 1, AcademicYear: "2025/2026"},
				{StudentID: createdStudents[3].ID, Subject: "Sejarah", Score: 82.0, Semester: 1, AcademicYear: "2025/2026"},
				{StudentID: createdStudents[4].ID, Subject: "Bahasa Inggris", Score: 95.0, Semester: 1, AcademicYear: "2025/2026"},
			}
			for i := range grades {
				DB.Create(&grades[i])
			}
		}
	}
}
