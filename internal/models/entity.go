package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONMetadata map[string]interface{}

func (j JSONMetadata) Value() (driver.Value, error) {
	if j == nil {
		return "{}", nil
	}
	return json.Marshal(j)
}

func (j *JSONMetadata) Scan(value interface{}) error {
	if value == nil {
		*j = make(map[string]interface{})
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New("unsupported type for JSONMetadata")
	}
	if len(bytes) == 0 {
		*j = make(map[string]interface{})
		return nil
	}
	return json.Unmarshal(bytes, j)
}

type EntityCategory string

const (
	CategoryVehicle   EntityCategory = "vehicle"
	CategoryIOTDevice EntityCategory = "iot_device"
	CategoryFacility  EntityCategory = "facility"
)

type EntityStatus string

const (
	StatusActive      EntityStatus = "active"
	StatusInactive    EntityStatus = "inactive"
	StatusMaintenance EntityStatus = "maintenance"
)

type Entity struct {
	BaseModel
	Name        string         `gorm:"size:150;not null" json:"name"`
	Code        string         `gorm:"size:50;not null;uniqueIndex" json:"code"`
	Category    EntityCategory `gorm:"size:30;not null;index" json:"category"`
	Status      EntityStatus   `gorm:"size:30;not null;default:'active';index" json:"status"`
	Latitude    float64        `gorm:"type:numeric(10,8);not null;index" json:"latitude"`
	Longitude   float64        `gorm:"type:numeric(11,8);not null;index" json:"longitude"`
	Description string         `gorm:"type:text" json:"description"`
	Metadata    JSONMetadata   `gorm:"type:jsonb;default:'{}'" json:"metadata"`
}

func (Entity) TableName() string {
	return "entities"
}
