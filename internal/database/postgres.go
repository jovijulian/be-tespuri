package database

import (
	"log"
	"time"

	"api-puri/internal/config"
	"api-puri/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectPostgres(cfg *config.Config) *gorm.DB {
	if cfg.DatabaseURL == "" {
		log.Fatalf("DATABASE_URL is not configured in environment or .env")
	}

	var gormLogLevel = logger.Warn
	if cfg.DBDebug {
		gormLogLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to retrieve generic database pool: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(15 * time.Minute)

	log.Println("PostgreSQL connection established successfully")
	DB = db
	return db
}

func AutoMigrate(db *gorm.DB) error {
	log.Println("Applying database migrations for entity models...")
	if err := db.AutoMigrate(&models.Entity{}); err != nil {
		return err
	}

	// Add spatial coordinates check constraints if not present
	_ = db.Exec("ALTER TABLE entities DROP CONSTRAINT IF EXISTS chk_entities_latitude;").Error
	_ = db.Exec("ALTER TABLE entities ADD CONSTRAINT chk_entities_latitude CHECK (latitude >= -90.0 AND latitude <= 90.0);").Error

	_ = db.Exec("ALTER TABLE entities DROP CONSTRAINT IF EXISTS chk_entities_longitude;").Error
	_ = db.Exec("ALTER TABLE entities ADD CONSTRAINT chk_entities_longitude CHECK (longitude >= -180.0 AND longitude <= 180.0);").Error

	_ = db.Exec("ALTER TABLE entities DROP CONSTRAINT IF EXISTS chk_entities_category;").Error
	_ = db.Exec("ALTER TABLE entities ADD CONSTRAINT chk_entities_category CHECK (category IN ('vehicle', 'iot_device', 'facility'));").Error

	_ = db.Exec("ALTER TABLE entities DROP CONSTRAINT IF EXISTS chk_entities_status;").Error
	_ = db.Exec("ALTER TABLE entities ADD CONSTRAINT chk_entities_status CHECK (status IN ('active', 'inactive', 'maintenance'));").Error

	// Add compound index for bounding box queries
	_ = db.Exec("CREATE INDEX IF NOT EXISTS idx_entities_coords ON entities (latitude, longitude);").Error

	log.Println("Database migration and spatial constraints applied successfully")
	return nil
}

func ResetDatabase(db *gorm.DB) error {
	log.Println("Dropping entities table for clean reset...")
	if err := db.Migrator().DropTable(&models.Entity{}); err != nil {
		return err
	}
	if err := AutoMigrate(db); err != nil {
		return err
	}
	return SeedDatabase(db)
}
