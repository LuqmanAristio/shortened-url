package models

import (
    "time"
    "gorm.io/gorm"
)

type URLShortener struct {
    ID          uint           `gorm:"primaryKey" json:"id"`           // ID unik
    OriginalURL string         `gorm:"type:text;not null" json:"original_url"` // URL asli
    ShortURL    string         `gorm:"size:255;not null;unique" json:"short_url"` // URL singkat (kode pendek)
    CreatedAt   time.Time      `json:"created_at"`                      // Waktu pembuatan URL
    ExpiresAt   *time.Time     `json:"expires_at,omitempty"`            // Waktu kadaluarsa (opsional)
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"` // Soft delete (opsional)
}
