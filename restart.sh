#!/bin/bash

set -e

echo "Memulai proses update aplikasi..."

# 1. Build Frontend
echo "=> [1/3] Memproses Frontend..."
cd frontend
npm run build
cd ..

# 2. Build Backend
echo "=> [2/3] Memproses Backend..."
cd backend
go build
cd ..

# 3. Restart PM2
echo "=> [3/3] Merestart PM2..."
pm2 restart all

echo "✅ Selesai! Semua service telah berhasil dibuild dan direstart."