package repository

import (
	"context"
	"firdausyusofs/kem_digital/internal/inventory"
	"firdausyusofs/kem_digital/internal/models"

	"gorm.io/gorm"
)

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) inventory.Repository {
	return &inventoryRepository{
		db: db,
	}
}

func (r *inventoryRepository) GetProducts(ctx context.Context) ([]*models.Product, error) {
	users := make([]*models.Product, 0)
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *inventoryRepository) GetProductByID(ctx context.Context, id int64) (*models.Product, error) {
	var product *models.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *inventoryRepository) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	result := r.db.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *inventoryRepository) UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	result := r.db.Model(&product).Updates(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *inventoryRepository) DeleteProduct(ctx context.Context, id int64) error {
	result := r.db.Delete(&models.Product{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
