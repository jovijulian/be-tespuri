# Entity Management API (api-puri)

Backend service for managing spatial entities (IoT devices, fleet vehicles, and infrastructure facilities) with geographic coordinate tracking and operational status management.

Built with **Go (Golang)**, **Fiber v2**, and **PostgreSQL (GORM)** using Clean Layered Architecture.

---

## 1. System Architecture & Patterns

The service adheres strictly to the modular layered architecture pattern:

```
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

---

## 2. Setup & Execution Guide (Local Machine)

### Prerequisites
- **Go**: Version 1.22 or newer
- **PostgreSQL**: Version 14 or newer running locally or via container

---

### Step 1: Create Local Database
Open your terminal or database management tool (psql, DBeaver, pgAdmin) and create an empty database named `tespuri`:

Using `psql`:
```bash
# Connect to your local PostgreSQL
psql -U postgres

# Create database
CREATE DATABASE tespuri;

# Exit psql
\q
```

Or via command line directly:
```bash
createdb -U postgres tespuri
```

---

### Step 2: Configure Environment Variables
Copy the `.env.example` file to `.env`:

```bash
cp .env.example .env
```

Adjust the `DATABASE_URL` in `.env` to match your local PostgreSQL credentials:

```env
APP_ENV=development
APP_PORT=3011
DEFAULT_LANGUAGE=en

# Format: postgres://<username>:<password>@<host>:<port>/<database_name>?sslmode=disable
DATABASE_URL=postgres://postgres:postgres@localhost:5432/tespuri?sslmode=disable

AUTO_MIGRATE=true
SEED_DB=false
DB_DEBUG=false
```

---

### Step 3: Run Database Migrations
Execute the automated migration flag to create tables, indexes, and spatial check constraints:

```bash
go run ./cmd/api --migrate
```
*Or using Makefile:*
```bash
make migrate
```

This will automatically create:
- Table `entities` with `id`, `name`, `code`, `category`, `status`, `latitude`, `longitude`, `description`, `metadata` (JSONB), timestamps, and soft delete.
- Spatial coordinate check constraints:
  - `chk_entities_latitude`: `-90.0 <= latitude <= 90.0`
  - `chk_entities_longitude`: `-180.0 <= longitude <= 180.0`
- Category & status enum constraints.
- Compound coordinate index on `(latitude, longitude)` for map acceleration.

---

### Step 4: Seed Baseline Spatial Data
Populate the database with 12 realistic spatial entities across Bandung and Jakarta (fleet vehicles, IoT stations, and municipal facilities):

```bash
go run ./cmd/api --seed
```
*Or using Makefile:*
```bash
make seed
```

---

### Step 5: Start the API Server
Start the service:

```bash
go run ./cmd/api
```
*Or using Makefile:*
```bash
make run
```

The API will start listening on `http://127.0.0.1:3011`.

To verify the service is running, run:
```bash
curl http://127.0.0.1:3011/health
```

Expected response:
```json
{
  "success": true,
  "message": "Service is healthy and ready",
  "data": {
    "service": "api-puri",
    "status": "healthy",
    "time": "2026-11-02T15:52:25Z"
  }
}
```

---

### Step 6: Run Automated Unit Tests
To verify all business logic, spatial constraints, and payload validations without needing an active PostgreSQL connection (using embedded in-memory SQLite):

```bash
go test -v ./...
```
*Or using Makefile:*
```bash
make test
```

---

## 3. Database Specification

### Table: `entities`

Stores all spatial assets with coordinates, category classification, operational state, and dynamic IoT metadata.

| Column | Type | Constraints | Description |
| :--- | :--- | :--- | :--- |
| `id` | `BIGSERIAL` | `PRIMARY KEY` | Auto-incrementing identifier |
| `name` | `VARCHAR(150)` | `NOT NULL` | Human-readable entity name |
| `code` | `VARCHAR(50)` | `NOT NULL, UNIQUE` | Unique tracking code (e.g., `IOT-ENV-001`, `VEH-BDG-101`) |
| `category` | `VARCHAR(30)` | `NOT NULL, INDEX` | Enum: `'vehicle'`, `'iot_device'`, `'facility'` |
| `status` | `VARCHAR(30)` | `NOT NULL, INDEX, DEFAULT 'active'` | Enum: `'active'`, `'inactive'`, `'maintenance'` |
| `latitude` | `DECIMAL(10, 8)` | `NOT NULL` | Latitude coordinate between `-90.00000000` and `90.00000000` |
| `longitude` | `DECIMAL(11, 8)` | `NOT NULL` | Longitude coordinate between `-180.00000000` and `180.00000000` |
| `description` | `TEXT` | `NULLABLE` | Detailed operational remarks |
| `metadata` | `JSONB` | `DEFAULT '{}'` | Extensible attributes (battery level, speed, temperature, etc.) |
| `created_at` | `TIMESTAMPTZ` | `NOT NULL, DEFAULT NOW()` | Record creation timestamp |
| `updated_at` | `TIMESTAMPTZ` | `NOT NULL, DEFAULT NOW()` | Last update timestamp |
| `deleted_at` | `TIMESTAMPTZ` | `INDEX, NULLABLE` | Soft delete support |

---

## 4. API Endpoints Reference

All API responses follow the standard JSON response envelope:

```json
{
  "success": true,
  "message": "Resource fetched successfully",
  "data": {},
  "errors": null,
  "meta": {
    "total": 12,
    "page": 1,
    "per_page": 20,
    "total_pages": 1,
    "server_time": "2026-11-02T15:30:00Z"
  }
}
```

### Endpoints Overview

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/health` | Health check probe |
| `GET` | `/api/v1/entities` | List entities with pagination, keyword search, and filters |
| `GET` | `/api/v1/entities/map` | Lightweight coordinates dataset optimized for map rendering |
| `GET` | `/api/v1/entities/:id` | Get entity detail by ID |
| `POST` | `/api/v1/entities` | Create new entity with coordinates validation |
| `PUT` | `/api/v1/entities/:id` | Update entity metadata and geographic location |
| `DELETE` | `/api/v1/entities/:id` | Soft delete entity |

---

### Detailed Endpoints

#### 1. List Entities
- **URL**: `GET /api/v1/entities`
- **Query Parameters**:
  - `page` (int, default: `1`)
  - `per_page` (int, default: `20`, max: `100`)
  - `search` (string, optional: matches `name` or `code`)
  - `category` (string, optional: `vehicle` \| `iot_device` \| `facility`)
  - `status` (string, optional: `active` \| `inactive` \| `maintenance`)
- **Example Request**:
  ```bash
  curl "http://127.0.0.1:3011/api/v1/entities?category=vehicle&status=active"
  ```

---

#### 2. Get Map Markers
- **URL**: `GET /api/v1/entities/map`
- **Query Parameters**:
  - `category` (optional)
  - `status` (optional)
- **Example Request**:
  ```bash
  curl "http://127.0.0.1:3011/api/v1/entities/map"
  ```

---

#### 3. Create Entity
- **URL**: `POST /api/v1/entities`
- **Content-Type**: `application/json`
- **Example Request**:
  ```bash
  curl -X POST http://127.0.0.1:3011/api/v1/entities \
    -H "Content-Type: application/json" \
    -d '{
      "name": "Air Quality Station Antapani",
      "code": "IOT-AQ-202",
      "category": "iot_device",
      "status": "active",
      "latitude": -6.915000,
      "longitude": 107.658000,
      "description": "Suburban air monitoring node",
      "metadata": {
        "pm25": 18.2,
        "battery_pct": 98
      }
    }'
  ```

---

#### 4. Update Entity
- **URL**: `PUT /api/v1/entities/:id`
- **Example Request**:
  ```bash
  curl -X PUT http://127.0.0.1:3011/api/v1/entities/1 \
    -H "Content-Type: application/json" \
    -d '{
      "name": "Logistics Fleet Van 01 (Rerouted)",
      "category": "vehicle",
      "status": "active",
      "latitude": -6.920000,
      "longitude": 107.625000,
      "description": "Updated vehicle route location",
      "metadata": {
        "speed_kmh": 42
      }
    }'
  ```

---

#### 5. Delete Entity
- **URL**: `DELETE /api/v1/entities/:id`
- **Example Request**:
  ```bash
  curl -X DELETE http://127.0.0.1:3011/api/v1/entities/1
  ```

---

## 5. Input Validation Strategy

- **Backend (Defensive Barrier)**:
  - Strongly typed DTO structs with `go-playground/validator/v10`.
  - Latitude validated against range `[-90.0, 90.0]`.
  - Longitude validated against range `[-180.0, 180.0]`.
  - Category restricted to `vehicle`, `iot_device`, `facility`.
  - Status restricted to `active`, `inactive`, `maintenance`.
  - Code uniqueness enforced both at application service level and database unique index level.
- **Frontend (Pre-flight Validation)**:
  - Form validation via Zod before sending HTTP requests, preventing invalid network payloads and giving immediate UI feedback.

---

## 6. Library Selection & Justification

| Library | Purpose | Justification |
| :--- | :--- | :--- |
| **Fiber v2** (`github.com/gofiber/fiber/v2`) | HTTP Web Framework | Built on top of Fasthttp, offering zero memory allocation for hot paths, high throughput, and clean middleware composition. |
| **GORM** (`gorm.io/gorm` & `gorm.io/driver/postgres`) | ORM & Query Builder | Provides robust schema migration, relation handling, and seamless PostgreSQL JSONB serialization/deserialization for flexible IoT metadata. |
| **Validator v10** (`github.com/go-playground/validator/v10`) | Struct Validation | Industry-standard declarative validation ensuring coordinate bounds and enum integrity before reaching database layers. |
| **godotenv** (`github.com/joho/godotenv`) | Configuration | Enables 12-factor configuration management across local, CI, and production environments. |

---

## 7. Agentic AI Workflow Documentation (`AGENT.md` Reference)

In accordance with the assessment criteria:

- **Scaffolding & Architecture Alignment**:
  AI was used as an engineering accelerator to audit existing enterprise patterns in `api-reeval` and scaffold the layered repository/service structure cleanly for spatial entities without redundant boilerplate.
- **Contract & Validation Design**:
  AI assisted in establishing dual-layer validation boundaries (Zod on frontend, Validator v10 on backend) and configuring coordinate precision (`DECIMAL(10,8)` / `DECIMAL(11,8)`).
- **Realistic Spatial Seeding**:
  AI generated realistic Indonesian geographical coordinates (Bandung and Jakarta IoT telemetry nodes, logistics fleet vehicles, municipal substations) to simulate production payloads.
- **Review & Verification**:
  Architectural trade-offs, code hygiene, and security posture were verified manually to maintain human engineering standards, pragmatic error handling, and clean documentation.
