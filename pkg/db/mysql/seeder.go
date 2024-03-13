package mysql

import (
	"context"
	"firdausyusofs/kem_digital/internal/inventory"
	inventoryRepo "firdausyusofs/kem_digital/internal/inventory/repository"
	"firdausyusofs/kem_digital/internal/models"
	"firdausyusofs/kem_digital/internal/supplier"
	supplierRepo "firdausyusofs/kem_digital/internal/supplier/repository"
	"firdausyusofs/kem_digital/pkg/logger"
	"strconv"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

type Seeder struct {
	db     *gorm.DB
	logger logger.Logger
}

func NewSeeder(db *gorm.DB, logger logger.Logger) *Seeder {
	return &Seeder{
		db:     db,
		logger: logger,
	}
}

func (s *Seeder) Run() error {
	// Init repositories
	inventoryRepo := inventoryRepo.NewInventoryRepository(s.db)
	supplierRepo := supplierRepo.NewSupplierRepository(s.db)

	// Seed suppliers
	if err := s.seedSuppliers(supplierRepo); err != nil {
		return err
	}

	// Seed products
	if err := s.seedProducts(inventoryRepo); err != nil {
		return err
	}

	return nil
}

func (s *Seeder) seedSuppliers(repo supplier.Repository) error {
	for i := 0; i < 10; i++ {
		s.logger.Info("Creating supplier ", i+1)

		supplier := &models.Supplier{
			Name: "Supplier " + strconv.Itoa(i+1),
		}

		createSupplier, err := repo.CreateSupplier(context.Background(), supplier)
		if err != nil {
			return err
		}

		s.logger.Info("Supplier created: ID ", createSupplier.ID)
	}

	return nil
}

func (s *Seeder) seedProducts(repo inventory.Repository) error {
	// Loop through 1000 times to create 1000 products
	for i := 0; i < 1000; i++ {
		s.logger.Info("Creating product ", i+1)

		randomInt, err := faker.RandomInt(1, 10)
		if err != nil {
			return err
		}

		product := &models.Product{
			Name: "Product " + strconv.Itoa(i+1),
			// Random from 1 to 10
			SupplierID: uint(randomInt[0]),
		}
		// Use faker to generate random data
		if err := faker.FakeData(&product); err != nil {
			return err
		}

		// Save product to database
		createdProduct, err := repo.CreateProduct(context.Background(), product)
		if err != nil {
			return err
		}

		s.logger.Info("Product created: ID ", createdProduct.ID)
	}

	return nil
}
