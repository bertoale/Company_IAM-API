# User & IAM Backend Service

## Gambaran Umum

Project ini adalah backend service untuk sistem User & Identity Access Management (IAM) berbasis REST API menggunakan bahasa Go (Golang) dan framework Gin. Sistem ini dirancang untuk mengelola user, role, permission, aplikasi, serta relasi-relasinya secara modular dan scalable.

## Fitur Utama

- **Manajemen User**: Registrasi, login, update, dan penghapusan user.
- **Manajemen Role**: CRUD role dan pengaturan role pada user.
- **Manajemen Permission**: CRUD permission dan pengaturan permission pada role.
- **Manajemen Application**: CRUD aplikasi dan pengaturan aplikasi pada user.
- **Relasi User-Role, Role-Permission, User-Application**: Mendukung multi-role, multi-permission, dan multi-application untuk setiap user.
- **Autentikasi JWT**: Login menghasilkan JWT token yang berisi user, roles, permissions, dan applications.
- **Rate Limiter**: Pembatasan request per IP untuk keamanan API.
- **Error Handling & Response Standar**: Response API konsisten dan mudah diintegrasikan dengan frontend.

## Struktur Folder

- `cmd/` : Entry point aplikasi (main.go)
- `internal/` : Berisi modul utama (user, role, permission, application, relasi, dsb)
- `pkg/` : Berisi package utilitas (config, middlewares, response, validator)

## Teknologi

- **Golang** (Go 1.25+)
- **Gin** (Web Framework)
- **GORM** (ORM untuk database)
- **MySQL** (default, bisa disesuaikan)
- **JWT** (Autentikasi)

## Cara Menjalankan

1. Copy `.env.example` ke `.env` dan sesuaikan konfigurasi database & JWT.
2. Jalankan perintah:
   ```bash
   go mod tidy
   go run cmd/main.go
   ```
3. API berjalan di port sesuai konfigurasi (`localhost:5000` secara default).

## Catatan

- Pastikan database sudah tersedia dan kredensial sudah benar.
- Struktur tabel akan otomatis dimigrasi saat aplikasi dijalankan.
- Rate limiter dan upload file sudah diaktifkan secara default.

---
