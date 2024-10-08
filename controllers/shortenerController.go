package controllers

import (
	"crypto/md5"
    "encoding/base64"
    "github.com/gofiber/fiber/v2"
    "github.com/LuqmanAristio/shortened-url/config"
    "github.com/LuqmanAristio/shortened-url/models"
    "time"
	"os"
)

func CreateShortURL(c *fiber.Ctx) error {
    originalURL := c.FormValue("original_url")

    shortURL := generateShortCode(originalURL)

    urlShortener := models.URLShortener{
        OriginalURL: originalURL,
        ShortURL:    shortURL,
        CreatedAt:   time.Now(),
    }

    config.DB.Create(&urlShortener)

    return c.Status(fiber.StatusCreated).JSON(os.Getenv("CORS_LOCAL")+"/"+urlShortener.ShortURL)
}

func generateShortCode(originalURL string) string {
    hash := md5.Sum([]byte(originalURL))

    shortCode := base64.RawURLEncoding.EncodeToString(hash[:])

    return shortCode[:6]
}

func RedirectToOriginalURL(c *fiber.Ctx) error {
    shortURL := c.Params("shortURL")

    var urlShortener models.URLShortener
    result := config.DB.Where("short_url = ?", shortURL).First(&urlShortener)

    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "URL not found",
        })
    }

    return c.Redirect(urlShortener.OriginalURL, fiber.StatusMovedPermanently)
}
