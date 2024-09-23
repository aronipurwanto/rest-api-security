package services

import "rest-api-mod02-http-method/models"

// In-memory database
var items = []models.Item{
	{ID: 1, Name: "Laptop", Price: 15000},
	{ID: 2, Name: "Smartphone", Price: 7000},
}

// ItemService handles the business logic for items
type ItemService struct{}

// GetAllItems returns all items
func (s *ItemService) GetAllItems() []models.Item {
	return items
}

// GetItemByID returns an item by its ID
func (s *ItemService) GetItemByID(id int) (models.Item, bool) {
	for _, item := range items {
		if item.ID == id {
			return item, true
		}
	}
	return models.Item{}, false
}

// CreateItem adds a new item
func (s *ItemService) CreateItem(newItem models.Item) models.Item {
	newItem.ID = len(items) + 1
	items = append(items, newItem)
	return newItem
}

// UpdateItem updates an existing item
func (s *ItemService) UpdateItem(id int, updatedItem models.Item) (models.Item, bool) {
	for i, item := range items {
		if item.ID == id {
			updatedItem.ID = id
			items[i] = updatedItem
			return updatedItem, true
		}
	}
	return models.Item{}, false
}

// DeleteItem deletes an item by ID
func (s *ItemService) DeleteItem(id int) bool {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return true
		}
	}
	return false
}
