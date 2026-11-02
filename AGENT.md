# AGENT.md — Panduan Konteks & Workflow Agentic AI (`be-tespuri`)

Dokumen ini berisi panduan konteks arsitektur, aturan teknis, serta penjelasan sejauh mana **Agentic AI** digunakan dalam pengembangan backend aplikasi Spatial Entity Management.

---

## 1. Peran & Batasan Agentic AI pada Proyek Ini

Pada pengerjaan tes ini, Agentic AI digunakan sebagai asisten *pair-programming* dengan alur kerja berikut:

1. **Keputusan Teknis oleh Developer**:
   - Menentukan stack utama (Go, Fiber v2, GORM, PostgreSQL).
   - Menentukan struktur folder *Clean/Layered Architecture* (`handlers`, `services`, `repositories`, `models`, `dto`).
   - Menentukan skema database dan penggunaan `JSONB` untuk atribut dinamis (`metadata`).
2. **Eksekusi oleh Agentic AI**:
   - Membuat *scaffolding* repository, service, dan handler sesuai pola standar developer.
   - Menulis validasi input koordinat (`latitude` -90 s/d 90, `longitude` -180 s/d 180) serta validasi kategori (`vehicle`, `iot_device`, `facility`) dan status (`active`, `maintenance`, `inactive`).
   - Membuat unit test (`internal/services/entity_service_test.go` dan `pkg/validator/validator_test.go`) serta data *seeder* (`internal/database/seeder.go`).
3. **Verifikasi & Deployment oleh Developer**:
   - Menguji seluruh endpoint di lokal dan production menggunakan Postman.
   - Menyesuaikan konfigurasi `Dockerfile` dan environment CI/CD hingga live di `https://tes.reeval.tech`.

---

## 2. Struktur Direktori

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
└── pkg/
    ├── logger/                   # Structured application logging
    ├── response/                 # Standard JSON response structure
    ├── validator/                # Request struct & coordinate validation
    └── i18n/                     # Error message constants
```

---

## 3. Aturan Bisnis & Validasi Data

Setiap perubahan kode oleh developer maupun Agentic AI wajib mematuhi aturan berikut:
- **Nama & Kode**: `name` wajib diisi; `code` bersifat unik (jika kosong saat create, sistem dapat men-generate otomatis atau memvalidasi keunikannya).
- **Kategori Valid**: `vehicle`, `iot_device`, `facility`.
- **Status Valid**: `active`, `maintenance`, `inactive`.
- **Koordinat Geografis**:
  - `latitude` wajib berada di rentang `-90` hingga `90`.
  - `longitude` wajib berada di rentang `-180` hingga `180`.
- **Format Response**: Seluruh endpoint wajib menggunakan helper `pkg/response` agar struktur JSON konsisten dengan ekspektasi frontend.
