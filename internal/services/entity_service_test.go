package services

import (
	"testing"

	"api-puri/internal/dto"
	"api-puri/internal/models"
	"api-puri/internal/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect memory db: %v", err)
	}
	if err := db.AutoMigrate(&models.Entity{}); err != nil {
		t.Fatalf("failed to migrate memory db: %v", err)
	}
	return db
}

func TestEntityService_CRUD(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewEntityRepository(db)
	service := NewEntityService(repo)

	// 1. Create
	createReq := &dto.CreateEntityRequest{
		Name:      "Test Telemetry Gateway",
		Code:      "IOT-TEST-001",
		Category:  "iot_device",
		Status:    "active",
		Latitude:  -6.917464,
		Longitude: 107.619122,
		Metadata: map[string]interface{}{
			"protocol": "MQTT",
			"firmware": "1.0.0",
		},
	}
	res, err := service.Create(createReq)
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}
	if res.ID == 0 || res.Code != "IOT-TEST-001" {
		t.Fatalf("unexpected create response: %+v", res)
	}

	// 2. Duplicate Code should fail
	_, err = service.Create(createReq)
	if err != ErrEntityCodeExists {
		t.Fatalf("expected ErrEntityCodeExists, got %v", err)
	}

	// 3. Get Detail
	detail, err := service.GetByID(res.ID)
	if err != nil {
		t.Fatalf("get detail failed: %v", err)
	}
	if detail.Name != "Test Telemetry Gateway" {
		t.Fatalf("name mismatch: %s", detail.Name)
	}

	// 4. Update
	updateReq := &dto.UpdateEntityRequest{
		Name:      "Updated Gateway Name",
		Category:  "iot_device",
		Status:    "maintenance",
		Latitude:  -6.920000,
		Longitude: 107.620000,
	}
	updated, err := service.Update(res.ID, updateReq)
	if err != nil {
		t.Fatalf("update failed: %v", err)
	}
	if updated.Name != "Updated Gateway Name" || updated.Status != "maintenance" {
		t.Fatalf("unexpected updated fields: %+v", updated)
	}

	// 5. Map Points
	points, err := service.GetMapPoints("", "")
	if err != nil {
		t.Fatalf("get map points failed: %v", err)
	}
	if len(points) != 1 {
		t.Fatalf("expected 1 map point, got %d", len(points))
	}

	// 6. Delete
	if err := service.Delete(res.ID); err != nil {
		t.Fatalf("delete failed: %v", err)
	}
	_, err = service.GetByID(res.ID)
	if err != ErrEntityNotFound {
		t.Fatalf("expected ErrEntityNotFound after delete, got %v", err)
	}
}
