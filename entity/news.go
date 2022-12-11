package entity

import (
	"gorm.io/gorm"
	"time"
)

type News struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Url         string    `json:"url"`
	ImageUrl    string    `json:"image_url"`
	Source      string    `json:"source"`
	PublishedAt time.Time `json:"published_at"`
	ApiSource   string    `json:"api_source"`
}
