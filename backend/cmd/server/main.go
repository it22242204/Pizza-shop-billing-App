package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/your‑gh‑user/pizza‑backend/internal/config"
    "github.com/your‑gh‑user/pizza‑backend/internal/database"
    "github.com/your‑gh‑user/pizza‑backend/internal/handlers"
    "github.com/your‑gh‑user/pizza‑backend/internal/models"
    "github.com/your‑gh‑user/pizza‑backend/internal/repositories"
    "github.com/your‑gh‑user/pizza‑backend/internal/services"
    "log"
)

func main() {
    cfg := config.MustLoad()

    // Connect & Auto‑migrate
    db := database.Connect(cfg.DBUrl,
        &models.Item{}, &models.Invoice{}, &models.InvoiceItem{})

    // Build layers
    itemRepo := repositories.NewItemRepository(db)
    itemSvc := services.NewItemService(itemRepo)
    itemHdl := handlers.NewItemHandler(itemSvc)

    // Fiber app
    app := fiber.New()
    app.Use(logger.New())

    api := app.Group("/api")
    itemHdl.Register(api.Group("/items"))
    // todo: invoice handler

    log.Printf("🌍  http://localhost:%s", cfg.Port)
    log.Fatal(app.Listen(":" + cfg.Port))
}
