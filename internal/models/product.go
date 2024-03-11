package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `json:"name" validate:"required"`
	Price       float64        `json:"price"`
	ImageURL    *string        `json:"image_url"`
	IsAvailable bool           `gorm:"default:true" json:"is_available"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type ProductList struct {
	TotalCount int        `json:"total_count"`
	TotalPages int        `json:"total_pages"`
	Page       int        `json:"page"`
	Limit      int        `json:"limit"`
	HasMore    bool       `json:"has_more"`
	Data       []*Product `json:"data"`
}
