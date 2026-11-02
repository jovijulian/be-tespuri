package routes

import (
	"time"

	"api-puri/internal/config"
	"api-puri/internal/handlers"
	"api-puri/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type RouteDependencies struct {
	Config        *config.Config
	EntityHandler *handlers.EntityHandler
}

func SetupRouter(deps *RouteDependencies) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "Entity Management API Engine",
		ServerHeader: "Entity-Engine",
		BodyLimit:    16 * 1024 * 1024,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(response.ApiResponse{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} | ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Accept-Language, X-Request-ID",
		ExposeHeaders:    "Content-Length, X-Request-ID",
		AllowCredentials: false,
		MaxAge:           86400,
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelDefault,
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return response.Success(c, "Service is healthy and ready", fiber.Map{
			"status":  "healthy",
			"service": "api-puri",
			"time":    time.Now().UTC().Format(time.RFC3339),
		})
	})

	api := app.Group("/api/v1")
	RegisterEntityRoutes(api, deps.EntityHandler)

	app.Use(func(c *fiber.Ctx) error {
		return response.NotFound(c, "Endpoint not found")
	})

	return app
}
