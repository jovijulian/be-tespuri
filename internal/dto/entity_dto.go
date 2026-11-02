package dto

import "time"

type CreateEntityRequest struct {
	Name        string                 `json:"name" validate:"required,min=3,max=150"`
	Code        string                 `json:"code" validate:"required,min=2,max=50"`
	Category    string                 `json:"category" validate:"required,oneof=vehicle iot_device facility"`
	Status      string                 `json:"status" validate:"required,oneof=active inactive maintenance"`
	Latitude    float64                `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude   float64                `json:"longitude" validate:"required,gte=-180,lte=180"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type UpdateEntityRequest struct {
	Name        string                 `json:"name" validate:"required,min=3,max=150"`
	Category    string                 `json:"category" validate:"required,oneof=vehicle iot_device facility"`
	Status      string                 `json:"status" validate:"required,oneof=active inactive maintenance"`
	Latitude    float64                `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude   float64                `json:"longitude" validate:"required,gte=-180,lte=180"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type EntityResponse struct {
	ID          uint                   `json:"id"`
	Name        string                 `json:"name"`
	Code        string                 `json:"code"`
	Category    string                 `json:"category"`
	Status      string                 `json:"status"`
	Latitude    float64                `json:"latitude"`
	Longitude   float64                `json:"longitude"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

type EntityMapPointDTO struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Category  string  `json:"category"`
	Status    string  `json:"status"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
