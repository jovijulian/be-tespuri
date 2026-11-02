package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api-puri/internal/config"
	"api-puri/internal/database"
	"api-puri/internal/handlers"
	"api-puri/internal/repositories"
	"api-puri/internal/routes"
	"api-puri/internal/services"
	"api-puri/pkg/logger"
)

func main() {
	cfg := config.LoadConfig()
	logger.Info("Initializing Entity Management API (Golang + Fiber v2 + PostgreSQL)...")

	db := database.ConnectPostgres(cfg)

	// CLI argument handling
	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			logger.Info("Executing database migrations...")
			if err := database.AutoMigrate(db); err != nil {
				logger.Error("Database migration failed: %v", err)
				os.Exit(1)
			}
			logger.Info("Database migration completed successfully!")
			os.Exit(0)
		}
		if arg == "--reset" {
			logger.Info("Resetting database schema and seeding fresh baseline data...")
			if err := database.ResetDatabase(db); err != nil {
				logger.Error("Database reset failed: %v", err)
				os.Exit(1)
			}
			logger.Info("Database reset and seeded successfully!")
			os.Exit(0)
		}
	}

	if cfg.AutoMigrate {
		if err := database.AutoMigrate(db); err != nil {
			logger.Error("Auto-migration failed: %v", err)
		}
	}

	shouldSeed := cfg.SeedDB
	for _, arg := range os.Args[1:] {
		if arg == "--seed" {
			shouldSeed = true
			break
		}
	}

	if shouldSeed {
		if err := database.SeedDatabase(db); err != nil {
			logger.Error("Database seeding failed: %v", err)
		}
	}

	// Dependency Injection Wiring
	entityRepo := repositories.NewEntityRepository(db)
	entityService := services.NewEntityService(entityRepo)
	entityHandler := handlers.NewEntityHandler(entityService)

	app := routes.SetupRouter(&routes.RouteDependencies{
		Config:        cfg,
		EntityHandler: entityHandler,
	})

	serverShutdown := make(chan struct{})
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan

		logger.Info("Received shutdown signal. Gracefully stopping server...")
		if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
			logger.Error("Error during graceful shutdown: %v", err)
		}
		close(serverShutdown)
	}()

	listenAddr := fmt.Sprintf(":%s", cfg.AppPort)
	logger.Info("Entity API server listening on %s", listenAddr)
	if err := app.Listen(listenAddr); err != nil {
		logger.Info("Server stopped: %v", err)
	}

	<-serverShutdown
	logger.Info("Server cleanup finished. Goodbye.")
}
