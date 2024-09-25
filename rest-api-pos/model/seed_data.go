package model

import "rest-api-pos/config"

func SeedData() {
	// Sample Category
	categories := []Category{
		{Name: "Electronics"},
		{Name: "Groceries"},
		{Name: "Clothing"},
		{Name: "Books"},
		{Name: "Stationery"},
		{Name: "Furniture"},
		{Name: "Toys"},
		{Name: "Beauty"},
		{Name: "Sports"},
		{Name: "Accessories"},
	}

	for _, category := range categories {
		config.DB.FirstOrCreate(&category, Category{Name: category.Name})
	}

	// Sample Supplier
	suppliers := []Supplier{
		{Name: "Supplier A", Address: "Jakarta"},
		{Name: "Supplier B", Address: "Bandung"},
		{Name: "Supplier C", Address: "Surabaya"},
		{Name: "Supplier D", Address: "Semarang"},
		{Name: "Supplier E", Address: "Bali"},
		{Name: "Supplier F", Address: "Medan"},
		{Name: "Supplier G", Address: "Yogyakarta"},
		{Name: "Supplier H", Address: "Malang"},
		{Name: "Supplier I", Address: "Makassar"},
		{Name: "Supplier J", Address: "Palembang"},
	}

	for _, supplier := range suppliers {
		config.DB.FirstOrCreate(&supplier, Supplier{Name: supplier.Name})
	}

	// Sample Product
	products := []Product{
		{Name: "Laptop", Price: 10000, Stock: 10, CategoryID: 1, SupplierID: 1},
		{Name: "Smartphone", Price: 5000, Stock: 20, CategoryID: 1, SupplierID: 2},
		{Name: "TV", Price: 8000, Stock: 15, CategoryID: 1, SupplierID: 3},
		{Name: "Refrigerator", Price: 12000, Stock: 5, CategoryID: 1, SupplierID: 4},
		{Name: "Microwave", Price: 3000, Stock: 7, CategoryID: 1, SupplierID: 5},
		{Name: "Rice Cooker", Price: 1500, Stock: 50, CategoryID: 2, SupplierID: 6},
		{Name: "Jeans", Price: 400, Stock: 100, CategoryID: 3, SupplierID: 7},
		{Name: "Shirt", Price: 250, Stock: 200, CategoryID: 3, SupplierID: 8},
		{Name: "Notebook", Price: 50, Stock: 500, CategoryID: 5, SupplierID: 9},
		{Name: "Table", Price: 3000, Stock: 10, CategoryID: 6, SupplierID: 10},
		{Name: "Chair", Price: 1200, Stock: 25, CategoryID: 6, SupplierID: 10},
		{Name: "Bed", Price: 5000, Stock: 3, CategoryID: 6, SupplierID: 9},
		{Name: "Basketball", Price: 800, Stock: 50, CategoryID: 9, SupplierID: 8},
		{Name: "Lipstick", Price: 200, Stock: 150, CategoryID: 8, SupplierID: 7},
		{Name: "Shampoo", Price: 100, Stock: 300, CategoryID: 8, SupplierID: 6},
		{Name: "Book", Price: 100, Stock: 1000, CategoryID: 4, SupplierID: 5},
		{Name: "Toy Car", Price: 300, Stock: 200, CategoryID: 7, SupplierID: 4},
		{Name: "Action Figure", Price: 500, Stock: 100, CategoryID: 7, SupplierID: 3},
		{Name: "Pen", Price: 20, Stock: 1000, CategoryID: 5, SupplierID: 2},
		{Name: "Monitor", Price: 2000, Stock: 25, CategoryID: 1, SupplierID: 1},
	}

	for _, product := range products {
		config.DB.FirstOrCreate(&product, Product{Name: product.Name})
	}
}
