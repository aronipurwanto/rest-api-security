package repository

import (
	"gorm.io/gorm"
	"rest-api-pos/model"
)

type ProductRepository interface {
	GetAllProducts() ([]model.Product, error)
	GetProductByID(id uint) (*model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Preload("Category").Preload("Supplier").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := r.db.Preload("Category").Preload("Supplier").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) CreateProduct(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) UpdateProduct(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}
