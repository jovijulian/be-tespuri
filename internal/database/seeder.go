package database

import (
	"log"

	"api-puri/internal/models"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	log.Println("Seeding realistic spatial entities (IoT, vehicles, facilities)...")

	seedEntities := []models.Entity{
		// Vehicles
		{
			Name:        "Logistics Fleet Van 01",
			Code:        "VEH-BDG-101",
			Category:    models.CategoryVehicle,
			Status:      models.StatusActive,
			Latitude:    -6.917464,
			Longitude:   107.619122,
			Description: "Main distribution courier van covering Bandung central district",
			Metadata: models.JSONMetadata{
				"plate_number": "D 1024 AA",
				"speed_kmh":    38.5,
				"battery_pct":  95,
				"fuel_level":   80,
			},
		},
		{
			Name:        "Intercity Heavy Truck 04",
			Code:        "VEH-JKT-404",
			Category:    models.CategoryVehicle,
			Status:      models.StatusActive,
			Latitude:    -6.208763,
			Longitude:   106.845599,
			Description: "Freight hauler currently on Jakarta - Cikampek logistics corridor",
			Metadata: models.JSONMetadata{
				"plate_number": "B 9876 XYZ",
				"speed_kmh":    62.0,
				"payload_ton":  12.4,
			},
		},
		{
			Name:        "Field Support Patrol 02",
			Code:        "VEH-BDG-202",
			Category:    models.CategoryVehicle,
			Status:      models.StatusMaintenance,
			Latitude:    -6.891480,
			Longitude:   107.610659,
			Description: "Field maintenance patrol vehicle undergoing engine tune-up",
			Metadata: models.JSONMetadata{
				"plate_number":     "D 4512 EF",
				"workshop":         "Dago Central Garage",
				"estimated_finish": "2026-11-03",
			},
		},

		// IoT Devices
		{
			Name:        "Air Quality Station Dago Atas",
			Code:        "IOT-AQ-101",
			Category:    models.CategoryIOTDevice,
			Status:      models.StatusActive,
			Latitude:    -6.879480,
			Longitude:   107.615200,
			Description: "Telemetry node monitoring ambient air index, PM2.5, and humidity",
			Metadata: models.JSONMetadata{
				"sensor_type": "Laser Particle AQI",
				"pm25":        22.4,
				"temperature": 23.8,
				"firmware":    "v2.1.0",
			},
		},
		{
			Name:        "Weather Station Gedung Sate",
			Code:        "IOT-WS-201",
			Category:    models.CategoryIOTDevice,
			Status:      models.StatusActive,
			Latitude:    -6.902481,
			Longitude:   107.618765,
			Description: "High-precision microclimate telemetry monitoring wind and precipitation",
			Metadata: models.JSONMetadata{
				"rainfall_mm": 0.0,
				"humidity":    78,
				"wind_speed":  12.1,
			},
		},
		{
			Name:        "Water Level Sensor Citarum Dayeuhkolot",
			Code:        "IOT-WL-301",
			Category:    models.CategoryIOTDevice,
			Status:      models.StatusActive,
			Latitude:    -6.985612,
			Longitude:   107.625341,
			Description: "River surge sensor for automated flood early warning",
			Metadata: models.JSONMetadata{
				"water_level_m": 3.42,
				"alert_level":   "NORMAL",
				"telemetry":     "LoRaWAN",
			},
		},
		{
			Name:        "Smart Solar Inverter Sensor Bandung",
			Code:        "IOT-PV-401",
			Category:    models.CategoryIOTDevice,
			Status:      models.StatusMaintenance,
			Latitude:    -6.934412,
			Longitude:   107.604123,
			Description: "Rooftop solar generator telemetry node with inverter communication",
			Metadata: models.JSONMetadata{
				"power_kw":     0.0,
				"fault_code":   "E-INV-04",
				"technician":   "Agus Santoso",
			},
		},
		{
			Name:        "IoT Gateway Sudirman Hub",
			Code:        "IOT-GW-501",
			Category:    models.CategoryIOTDevice,
			Status:      models.StatusActive,
			Latitude:    -6.215500,
			Longitude:   106.820100,
			Description: "Edge aggregation gateway for regional smart street lighting nodes",
			Metadata: models.JSONMetadata{
				"connected_nodes": 48,
				"uptime_days":     124,
				"packet_loss_pct": 0.02,
			},
		},

		// Facilities
		{
			Name:        "PT LEN Industri HQ Facility",
			Code:        "FAC-LEN-001",
			Category:    models.CategoryFacility,
			Status:      models.StatusActive,
			Latitude:    -6.943142,
			Longitude:   107.608321,
			Description: "Corporate headquarters and R&D manufacturing engineering campus",
			Metadata: models.JSONMetadata{
				"address":     "Jl. Soekarno Hatta No.442, Bandung",
				"capacity":    1200,
				"type":        "Headquarters & Lab",
			},
		},
		{
			Name:        "Logistics Distribution Center Soekarno-Hatta",
			Code:        "FAC-DC-002",
			Category:    models.CategoryFacility,
			Status:      models.StatusActive,
			Latitude:    -6.948210,
			Longitude:   107.671230,
			Description: "Central warehouse and regional sorting fulfillment facility",
			Metadata: models.JSONMetadata{
				"dock_doors":    14,
				"storage_m2":    8500,
				"cold_storage":  true,
			},
		},
		{
			Name:        "Greater Jakarta Regional Hub",
			Code:        "FAC-HUB-003",
			Category:    models.CategoryFacility,
			Status:      models.StatusActive,
			Latitude:    -6.225500,
			Longitude:   106.800500,
			Description: "Regional commercial and service integration operations hub",
			Metadata: models.JSONMetadata{
				"metro_coverage": "Jabodetabek",
				"status":         "Operational",
			},
		},
		{
			Name:        "Depot & Workshop Kiaracondong",
			Code:        "FAC-DEP-004",
			Category:    models.CategoryFacility,
			Status:      models.StatusMaintenance,
			Latitude:    -6.924150,
			Longitude:   107.647890,
			Description: "Fleet maintenance workshop and spare parts depot",
			Metadata: models.JSONMetadata{
				"maintenance_bays": 6,
				"active_tickets":   4,
			},
		},
	}

	for _, item := range seedEntities {
		var existing models.Entity
		err := db.Where("code = ?", item.Code).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&item).Error; err != nil {
				log.Printf("Failed to seed entity %s: %v", item.Code, err)
			}
		}
	}

	log.Printf("Successfully verified and seeded %d spatial entities", len(seedEntities))
	return nil
}
