package repository

import (
	"context"
	"firdausyusofs/kem_digital/internal/inventory"
	"firdausyusofs/kem_digital/internal/models"
	"firdausyusofs/kem_digital/pkg/utils"

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

func (r *inventoryRepository) GetProducts(ctx context.Context, pq *utils.PaginationQuery) (*models.ProductList, error) {
	var totalCount int64
	// Get total count of products
	result := r.db.Model(&models.Product{}).Count(&totalCount)
	if result.Error != nil {
		return nil, result.Error
	}

	if totalCount == 0 {
		return &models.ProductList{
			TotalCount: int(totalCount),
			TotalPages: utils.GetTotalPage(int(totalCount), pq.GetLimit()),
			Page:       pq.GetPage(),
			Limit:      pq.GetLimit(),
			HasMore:    utils.CheckHasMore(int(totalCount), pq.GetLimit(), pq.GetPage()),
			Data:       make([]*models.Product, 0),
		}, nil
	}

	products := make([]*models.Product, 0)
	order := "id" + " " + pq.GetOrderBy()
	result = r.db.Limit(pq.GetLimit()).Offset(pq.GetOffset()).Order(order).Preload("Supplier").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.ProductList{
		TotalCount: int(totalCount),
		TotalPages: utils.GetTotalPage(int(totalCount), pq.GetLimit()),
		Page:       pq.GetPage(),
		Limit:      pq.GetLimit(),
		HasMore:    utils.CheckHasMore(int(totalCount), pq.GetLimit(), pq.GetPage()),
		Data:       products,
	}, nil
}

func (r *inventoryRepository) GetProductByID(ctx context.Context, id int64) (*models.Product, error) {
	var product *models.Product
	result := r.db.Preload("Supplier").First(&product, id)
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
