package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/xphp/go-url-shortener/internal/adapters/storage/postgres"
	"github.com/xphp/go-url-shortener/internal/adapters/storage/redis"
	"github.com/xphp/go-url-shortener/internal/api/handlers"
	"github.com/xphp/go-url-shortener/internal/api/routes"
	"github.com/xphp/go-url-shortener/internal/core/services"
	"github.com/xphp/go-url-shortener/pkg/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := postgres.NewConnection(cfg.Database.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := postgres.CreateSchema(db); err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}
	log.Println("Database schema created successfully")

	redisClient, err := redis.NewConnection(cfg.Redis.Address(), cfg.Redis.Password, cfg.Redis.DB)
	if err != nil {
		log.Printf("Warning: Failed to connect to Redis: %v", err)
		log.Println("Continuing without cache...")
	} else {
		log.Println("Connected to Redis successfully")
	}

	urlRepo := postgres.NewURLRepository(db)
	cacheRepo := redis.NewCacheRepository(redisClient)
	urlService := services.NewURLService(urlRepo, cacheRepo, cfg.App.ShortCodeLength, cfg.App.BaseURL)
	urlHandler := handlers.NewURLHandler(urlService)

	app := fiber.New(fiber.Config{
		AppName: "URL Shortener v1.0",
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("High-Performance URL Shortener API is running!")
	})

	routes.SetupRoutes(app, urlHandler)

	port := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on port %s\n", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
