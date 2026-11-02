# Spatial Entity Management — Backend API (`be-tespuri`)

Backend REST API untuk manajemen entitas spasial (Kendaraan, Perangkat IoT, dan Fasilitas) berbasis **Go (Golang)**, **Fiber v2**, dan **PostgreSQL**.

## Link Deployment & Repository

- **Live API Base URL**: https://tes.reeval.tech
- **Health Check**: https://tes.reeval.tech/health
- **Endpoint Entitas**: https://tes.reeval.tech/api/v1/entities
- **Repository Backend**: https://github.com/jovijulian/be-tespuri
- **Live Frontend App**: https://fe-tespuri.vercel.app/
- **Repository Frontend**: https://github.com/jovijulian/fe-tespuri

---

## Status Fitur (Functional Requirements)

Seluruh kebutuhan fungsional pada soal *Take-Home Test* telah selesai dikerjakan (tidak ada fitur yang belum selesai):

- [x] Menyimpan dan mengambil data entitas spasial melalui REST API.
- [x] Mendukung operasi CRUD penuh (Create, Read + Filter & Pagination, Update, Delete).
- [x] Menyimpan koordinat geografis (`latitude`, `longitude`) dan atribut dinamis (`metadata` JSONB).
- [x] Validasi input di sisi backend (batas koordinat geografis `-90..90` & `-180..180`, enum kategori & status, serta keunikan kode entitas).

---

## Struktur Folder & Arsitektur

Project ini menggunakan pendekatan **Clean / Layered Architecture** untuk memisahkan *HTTP transport*, *business logic*, dan *data access layer*:

```text
api-puri/
├── cmd/
│   └── api/
│       └── main.go               # App bootstrap, CLI flags (--migrate, --seed, --reset), graceful shutdown
├── internal/
│   ├── config/
│   │   └── config.go             # Environment configuration loader
│   ├── database/
│   │   ├── postgres.go           # GORM initialization, connection pooling, check constraints
│   │   └── seeder.go             # Baseline seed data (IoT, vehicles, facilities)
│   ├── models/
│   │   ├── base.go               # BaseModel (ID, CreatedAt, UpdatedAt, DeletedAt)
│   │   └── entity.go             # Entity GORM model definition & table schema
│   ├── dto/
│   │   └── entity_dto.go         # Request payload & Response envelopes with validation tags
│   ├── repositories/
│   │   ├── base_repository.go    # Generic repository interface & GORM implementation
│   │   └── entity_repository.go  # Entity-specific query methods (spatial, filters, pagination)
│   ├── services/
│   │   ├── entity_service.go     # Business logic, boundary validation, DTO mapping
│   │   └── entity_service_test.go# Unit tests with in-memory SQLite database
│   ├── handlers/
│   │   └── entity_handler.go     # HTTP request parsing, validator calls, response dispatch
│   └── routes/
│       ├── routes.go             # Fiber app configuration, CORS, logger, and middleware
│       └── entity_routes.go      # Route group definitions for /api/v1/entities
├── pkg/
│   ├── logger/
│   │   └── logger.go             # Structured application logging
│   ├── response/
│   │   └── response.go           # Standard JSON response structure (Success, Paginated, Errors)
│   ├── validator/
│   │   ├── validator.go          # go-playground/validator wrapper with English error formatting
│   │   └── validator_test.go     # Unit tests for coordinate bounds and enum validation
│   └── i18n/
│       └── i18n.go               # Localization helper & error translation constants
├── .air.toml                     # Live-reload configuration
├── .env.example                  # Environment configuration template
├── Dockerfile                    # Multi-stage production container build
├── Makefile                      # Developer shortcut commands
└── README.md                     # Technical documentation & API specification
```

### Penjelasan Tiap Layer:
- **`cmd/api`**: Entrypoint utama aplikasi yang menginisialisasi koneksi database, menangani argumen CLI (`--migrate`, `--seed`, `--reset`), merangkai *dependency injection*, serta menangani *graceful shutdown*.
- **`internal/config` & `internal/database`**: Mengelola pembacaan variabel `.env`, konfigurasi *connection pool* PostgreSQL, migrasi tabel otomatis (`AutoMigrate`), serta pengisian data awal (*seeder*).
- **`internal/models` & `internal/dto`**: Memisahkan skema tabel database (`models`) dari struktur payload JSON (`dto`) agar kontrak API tetap aman dan dilengkapi tag validasi.
- **`internal/repositories`**: Bertanggung jawab penuh atas query database menggunakan GORM (pencarian, filter kategori/status, filter *bounding-box* koordinat, dan pagination).
- **`internal/services`**: Menyimpan logika bisnis utama, pengecekan duplikasi kode entitas, pemetaan DTO ke model, beserta unit test terisolasi menggunakan SQLite *in-memory*.
- **`internal/handlers` & `internal/routes`**: Mengatur routing Fiber v2, middleware CORS/Recovery, parsing parameter HTTP, pemanggilan validator, dan pengembalian response JSON.
- **`pkg/`**: Kumpulan package utilitas internal yang bersifat *reusable* (`validator`, `response`, `logger`, `i18n`).

---

## Cara Menjalankan Program

### 1. Persiapan Environment
Salin file `.env.example` menjadi `.env` lalu sesuaikan konfigurasi database PostgreSQL:
```bash
cp .env.example .env
```

### 2. Migrasi Database & Seeder Data Awal
Gunakan perintah `go run` atau shortcut `Makefile`:
```bash
# Menjalankan migrasi tabel
go run ./cmd/api --migrate
# atau: make migrate

# Mengisi data awal (kendaraan, IoT, fasilitas)
go run ./cmd/api --seed
# atau: make seed

# Reset tabel + isi ulang data awal
go run ./cmd/api --reset
# atau: make reset
```

### 3. Menjalankan Server API
```bash
go run ./cmd/api
# atau: make run
```
Server berjalan di `http://localhost:3011` (sesuai konfigurasi `APP_PORT`).

### 4. Menjalankan Unit Test
```bash
go test -v ./...
# atau: make test
```

### 5. Menjalankan via Docker
```bash
docker build -t be-tespuri .
docker run -p 3011:3011 --env-file .env be-tespuri
```

---

## Daftar Endpoint API

| Method | Endpoint | Deskripsi |
| :--- | :--- | :--- |
| `GET` | `/health` | Cek status kesehatan server |
| `GET` | `/api/v1/entities` | Ambil daftar entitas (mendukung query `?search=`, `?category=`, `?status=`, `?page=`, `?limit=`) |
| `GET` | `/api/v1/entities/:id` | Ambil detail entitas berdasarkan UUID |
| `POST` | `/api/v1/entities` | Tambah entitas baru |
| `PUT` | `/api/v1/entities/:id` | Update data & lokasi entitas |
| `DELETE` | `/api/v1/entities/:id` | Hapus entitas (soft delete) |

File koleksi Postman tersedia pada `tespuri.postman_collection.json`.

---

## Alasan Pemilihan Library & Tech Stack

1. **Go + Fiber v2 (`github.com/gofiber/fiber/v2`)**
   Dipilih karena performanya cepat (berbasis Fasthttp), penggunaan memori rendah, serta memiliki *middleware* bawaan (CORS, Logger, Recover) yang memudahkan pembuatan REST API secara rapi.
2. **GORM (`gorm.io/gorm`) + PostgreSQL (`JSONB`)**
   PostgreSQL dipilih karena mendukung tipe data `JSONB`. Karena setiap jenis objek (Kendaraan, IoT, Fasilitas) memiliki atribut tambahan yang berbeda-beda (seperti plat nomor, level baterai, atau kapasitas gedung), kolom `JSONB` membuat penyimpanan atribut tetap fleksibel tanpa perlu mengubah struktur tabel.
3. **Go Playground Validator (`github.com/go-playground/validator/v10`)**
   Memastikan validasi input (seperti batas koordinat `latitude` `-90..90`, `longitude` `-180..180`, dan enum kategori/status) dilakukan secara deklaratif dan konsisten.

---

## Workflow Penggunaan Agentic AI

Pengerjaan proyek ini dibantu oleh **Agentic AI** sebagai *pair-programming partner* dengan pembagian peran berikut:

1. **Perancangan Arsitektur & Kontrak Data**: Developer menentukan struktur arsitektur *layered* (`handlers -> services -> repositories`), skema tabel, serta aturan validasi bisnis.
2. **Scaffolding & Implementasi CRUD**: AI membantu men-generate boilerplate kode CRUD, DTO request/response, seeder data spasial, serta unit test pada layer service dan validator.
3. **Review, Debugging & Deployment**: Developer melakukan verifikasi langsung via Postman, menyesuaikan konfigurasi `Dockerfile` (`GOTOOLCHAIN=auto`) pada pipeline CI/CD, serta men-deploy service ke server production (`https://tes.reeval.tech`).

Penjelasan konteks teknis untuk agen AI dapat dilihat pada file [AGENT.md](./AGENT.md).
