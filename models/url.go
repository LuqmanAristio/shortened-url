package models

import (
    "time"
    "gorm.io/gorm"
)

type URLShortener struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    OriginalURL string         `gorm:"type:text;not null" json:"original_url"`
    ShortURL    string         `gorm:"size:255;not null;unique" json:"short_url"`
    CreatedAt   time.Time      `json:"created_at"`
    ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
