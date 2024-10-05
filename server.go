package main

import (
    "log"
    "github.com/gofiber/fiber/v3"
    "github.com/LuqmanAristio/shortened-url/config"
)

func main() {
    config.Connect()
    
    app := fiber.New(fiber.Config{
        CaseSensitive: true,
        StrictRouting: true,
        ServerHeader:  "FiberServer",
        AppName: "Test App URL Shortener v1.0.1",
    })

    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Hello 👋!")
    })

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}