# Sistem Manajemen Kesiswaan (Student Management System)

Aplikasi Sistem Manajemen Kesiswaan sekolah menengah atas dengan arsitektur modern yang memisahkan backend (Golang + Gin + GORM + PostgreSQL) dan frontend (Nuxt.js 3 + Tailwind CSS).

Aplikasi ini didesain dengan visual bertema gelap premium berpola struktural, aksen emas-kuning tajam, serta tipografi editorial yang elegan sesuai dengan panduan desain.

---

## Persyaratan Awal (Prerequisites)
Sebelum menjalankan aplikasi, pastikan sistem Anda telah memiliki instalasi berikut:
1. **Golang** (minimal versi 1.22)
2. **Node.js** (minimal versi 22)
3. **npm** (minimal versi 10)
4. **PostgreSQL** (minimal versi 14)

---

## 1. Setup Database PostgreSQL
Aplikasi menggunakan GORM untuk *Auto Migration* skema database secara otomatis ketika backend pertama kali dijalankan.

Langkah persiapan database:
1. Masuk ke PostgreSQL server Anda:
   ```bash
   psql -U postgres
   ```
2. Buat database baru bernama `schoolapp`:
   ```sql
   CREATE DATABASE schoolapp;
   ```
3. Konfigurasi kredensial koneksi database Anda di berkas `/backend/.env` jika berbeda dari konfigurasi bawaan.

---

## 2. Setup dan Menjalankan Backend
Backend berjalan di port `8080` dan menyediakan API JSON standar yang dilindungi enkripsi JWT Token.

Langkah menjalankan backend:
1. Pindah ke direktori backend:
   ```bash
   cd backend
   ```
2. Salin berkas lingkungan konfigurasi `.env`:
   ```bash
   cp .env.example .env
   ```
3. Sesuaikan isi konfigurasi `.env` dengan kredensial database Anda.
4. Jalankan aplikasi:
   ```bash
   go run main.go
   ```
   *Catatan: Saat backend pertama kali berjalan, sistem akan otomatis bermigrasi dan menyemai (seed) data default akun Administrator (`admin@sekolah.com` / `admin123`) serta 10 data siswa contoh beserta riwayat absensi dan nilainya.*

---

## 3. Setup dan Menjalankan Frontend
Frontend berjalan di port `3000` menggunakan teknologi Nuxt.js 3 dan Tailwind CSS.

Langkah menjalankan frontend:
1. Pindah ke direktori frontend:
   ```bash
   cd frontend
   ```
2. Salin berkas lingkungan konfigurasi `.env`:
   ```bash
   cp .env.example .env
   ```
3. Pastikan konfigurasi `VITE_API_URL` mengarah ke URL server backend Anda (`http://localhost:8080`).
4. Instalasi dependensi:
   ```bash
   npm install
   ```
5. Jalankan server pengembang lokal:
   ```bash
   npm run dev
   ```
6. Buka peramban (browser) dan akses `http://localhost:3000`. Masuk menggunakan kredensial default admin di atas.

---

## Daftar Lengkap Endpoint API

Seluruh respon API dikemas dalam format JSON terstandarisasi:
```json
{
  "success": true,
  "data": null,
  "message": "Pesan status"
}
```

### Autentikasi (Terbuka Umum)
- **POST** `/api/auth/register` → Mendaftarkan akun pengguna baru (admin/guru).
- **POST** `/api/auth/login` → Autentikasi masuk pengguna, mengembalikan Token JWT (berlaku 24 jam).

### Manajemen Siswa (Dilindungi JWT)
- **GET** `/api/students` → Menampilkan daftar semua siswa secara berpaginasi (dilengkapi parameter pencarian `search` dan limit).
- **GET** `/api/students/:id` → Menampilkan detail profil lengkap seorang siswa berdasarkan ID (UUID).
- **POST** `/api/students` → Menambahkan catatan data siswa baru (memeriksa keunikan NIS).
- **PUT** `/api/students/:id` → Memperbarui rincian data siswa terdaftar.
- **DELETE** `/api/students/:id` → Menghapus data siswa dan seluruh dependensi riwayat nilai & absensinya secara transaksional di database.

### Manajemen Kehadiran / Absensi (Dilindungi JWT)
- **GET** `/api/attendances` → Menampilkan data rekap absensi harian (dapat difilter berdasarkan parameter query `date` dan `student_id`).
- **POST** `/api/attendances` → Mencatat kehadiran harian siswa (mendukung pembaharuan otomatis/upsert bila data siswa pada tanggal tersebut telah terekam).
- **PUT** `/api/attendances/:id` → Mengubah catatan kehadiran siswa.

### Manajemen Nilai Akademik (Dilindungi JWT)
- **GET** `/api/grades` → Menampilkan rekap nilai pelajaran siswa (dapat difilter berdasarkan parameter query `student_id`).
- **POST** `/api/grades` → Mencatat nilai pelajaran siswa (mendukung pembaharuan otomatis/upsert jika nilai mapel untuk siswa pada semester & tahun ajaran yang sama telah ada).
- **PUT** `/api/grades/:id` → Mengubah skor evaluasi nilai pelajaran siswa.

### Dashboard Utama (Dilindungi JWT)
- **GET** `/api/dashboard/stats` → Mengembalikan data ringkasan agregat statistik kesiswaan (total siswa terdaftar, persentase kehadiran bulan ini, rata-rata nilai akademik seluruh kelas, serta daftar 5 siswa terdaftar terbaru).
