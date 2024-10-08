package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/LuqmanAristio/shortened-url/config"
    "github.com/LuqmanAristio/shortened-url/models"
    "github.com/LuqmanAristio/shortened-url/routes"
    "github.com/LuqmanAristio/shortened-url/middleware"
)

func main() {
    app := fiber.New(fiber.Config{
        CaseSensitive: true,
        StrictRouting: true,
        ServerHeader:  "URL",
        AppName: "URL Shortener v1.0.1",
    })

    config.ConnectDB()
    config.DB.AutoMigrate(&models.URLShortener{})

    middleware.SetupCORS(app)
    routes.SetupRoutes(app)
    
    log.Fatal(app.Listen(":3000"))
}