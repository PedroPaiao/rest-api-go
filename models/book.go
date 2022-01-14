package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `json:"uid" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	MediumPrice string         `json:"medium_price"`
	Author      string         `json:"author"`
	ImageUrl    string         `json:"image_url"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
