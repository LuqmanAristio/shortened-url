package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/LuqmanAristio/shortened-url/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Static("/", "./public")
	app.Post("/shortener", controllers.CreateShortURL)
	app.Get("/:shortURL", controllers.RedirectToOriginalURL)
	
}