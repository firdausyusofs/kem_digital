package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primarykey" json:"id" faker:"-"`
	Name        string         `json:"name" validate:"required" faker:"-"`
	Description *string        `json:"description" faker:"paragraph"`
	Price       float64        `json:"price" faker:"amount"`
	IsAvailable bool           `gorm:"default:true" json:"is_available" faker:"-"`
	SupplierID  uint           `json:"supplier_id" faker:"-"`
	Supplier    *Supplier      `json:"supplier" faker:"-"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" faker:"-"`
}

type ProductList struct {
	TotalCount int        `json:"total_count"`
	TotalPages int        `json:"total_pages"`
	Page       int        `json:"page"`
	Limit      int        `json:"limit"`
	HasMore    bool       `json:"has_more"`
	Data       []*Product `json:"data"`
}
