package repository

import (
	"gorm.io/gorm"
	"rest-api-pos/model"
)

type CategoryRepository interface {
	GetAllCategories() ([]model.Category, error)
	GetCategoryByID(id uint) (*model.Category, error)
	CreateCategory(category *model.Category) error
	UpdateCategory(category *model.Category) error
	DeleteCategory(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) GetCategoryByID(id uint) (*model.Category, error) {
	var category model.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) CreateCategory(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) UpdateCategory(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) DeleteCategory(id uint) error {
	return r.db.Delete(&model.Category{}, id).Error
}
