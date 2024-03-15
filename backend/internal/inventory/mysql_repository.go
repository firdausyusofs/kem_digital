package inventory

import (
	"context"
	"firdausyusofs/kem_digital/internal/models"
	"firdausyusofs/kem_digital/pkg/utils"
)

type Repository interface {
	GetProducts(ctx context.Context, pq *utils.PaginationQuery) (*models.ProductList, error)
	GetProductByID(ctx context.Context, id int64) (*models.Product, error)
	CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
}
