package validator

import (
	"testing"
)

type sampleStruct struct {
	Name      string  `validate:"required,min=3"`
	Category  string  `validate:"required,oneof=vehicle iot_device facility"`
	Latitude  float64 `validate:"required,gte=-90,lte=90"`
	Longitude float64 `validate:"required,gte=-180,lte=180"`
}

func TestValidateStruct_Valid(t *testing.T) {
	data := sampleStruct{
		Name:      "Valid Sensor",
		Category:  "iot_device",
		Latitude:  -6.917464,
		Longitude: 107.619122,
	}

	errs := ValidateStruct(nil, &data)
	if len(errs) > 0 {
		t.Fatalf("expected 0 errors, got %d: %v", len(errs), errs)
	}
}

func TestValidateStruct_InvalidCoordinates(t *testing.T) {
	data := sampleStruct{
		Name:      "Bad Coordinates",
		Category:  "vehicle",
		Latitude:  95.0, // Invalid: > 90
		Longitude: 200.0, // Invalid: > 180
	}

	errs := ValidateStruct(nil, &data)
	if len(errs) != 2 {
		t.Fatalf("expected 2 coordinate errors, got %d", len(errs))
	}
}

func TestValidateStruct_InvalidCategory(t *testing.T) {
	data := sampleStruct{
		Name:      "Aircraft",
		Category:  "plane", // Invalid
		Latitude:  -6.2,
		Longitude: 106.8,
	}

	errs := ValidateStruct(nil, &data)
	if len(errs) != 1 {
		t.Fatalf("expected 1 category error, got %d", len(errs))
	}
}
