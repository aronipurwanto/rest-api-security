package repository

import (
	"gorm.io/gorm"
	"rest-api-pos/model"
)

type SupplierRepository interface {
	GetAllSuppliers() ([]model.Supplier, error)
	GetSupplierByID(id uint) (*model.Supplier, error)
	CreateSupplier(supplier *model.Supplier) error
	UpdateSupplier(supplier *model.Supplier) error
	DeleteSupplier(id uint) error
}

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepository{db}
}

func (r *supplierRepository) GetAllSuppliers() ([]model.Supplier, error) {
	var suppliers []model.Supplier
	if err := r.db.Find(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (r *supplierRepository) GetSupplierByID(id uint) (*model.Supplier, error) {
	var supplier model.Supplier
	if err := r.db.First(&supplier, id).Error; err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (r *supplierRepository) CreateSupplier(supplier *model.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *supplierRepository) UpdateSupplier(supplier *model.Supplier) error {
	return r.db.Save(supplier).Error
}

func (r *supplierRepository) DeleteSupplier(id uint) error {
	return r.db.Delete(&model.Supplier{}, id).Error
}
