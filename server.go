package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/LuqmanAristio/shortened-url/config"
    "github.com/LuqmanAristio/shortened-url/models"
)

func main() {
    app := fiber.New(fiber.Config{
        CaseSensitive: true,
        StrictRouting: true,
        ServerHeader:  "FiberServer",
        AppName: "Test App URL Shortener v1.0.1",
    })

    config.ConnectDB()
    config.DB.AutoMigrate(&models.URLShortener{})

    app.Static("/", "./public")
    
    log.Fatal(app.Listen(":3000"))
}