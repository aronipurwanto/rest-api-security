# POS API Testing

## Deskripsi
Proyek ini adalah sistem **Point of Sale (POS)** sederhana yang dibangun menggunakan **Golang** dan **Fiber**, dengan autentikasi melalui **Keycloak**.

File ini berisi semua endpoint untuk pengujian API dengan format **cURL** yang dapat digunakan untuk menguji semua fitur yang ada di API POS.

## Prasyarat

Sebelum memulai pengujian API, pastikan:
- **Keycloak** sudah terkonfigurasi dan berjalan.
- Token akses (**access_token**) dan refresh token (**refresh_token**) diperoleh setelah login ke Keycloak.
- **Golang** dan **Fiber** telah terinstall dengan benar.

## Endpoint API Testing Menggunakan cURL

### 1. Login (POST `/login`)

Gunakan endpoint ini untuk login dan mendapatkan access token serta refresh token.

```bash
curl -X POST http://localhost:3000/login \
-H "Content-Type: application/json" \
-d '{
  "username": "your_username",
  "password": "your_password"
}'
```

2. Refresh Token (POST /refresh-token)
   Gunakan refresh token untuk memperbarui access token.
```bash
curl -X POST http://localhost:3000/refresh-token \
-H "Content-Type: application/json" \
-d '{
"refresh_token": "<your_refresh_token>"
}'
```

3. Get Profile (GET /profile)
   Gunakan access token untuk mendapatkan profil pengguna yang login.
```bash
curl -X GET http://localhost:3000/profile \
-H "Authorization: Bearer <your_access_token>"
```
4. Suppliers

Get All Suppliers (GET /suppliers)
```bash
curl -X GET http://localhost:3000/suppliers \
-H "Authorization: Bearer <your_access_token>" 
```

Get Supplier by ID (GET /suppliers/:id)
```bash
curl -X GET http://localhost:3000/suppliers/1 \
-H "Authorization: Bearer <your_access_token>"
```

Create Supplier (POST /suppliers)
```bash
curl -X POST http://localhost:3000/suppliers \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_access_token>" \
-d '{
  "name": "New Supplier",
  "address": "123 Supplier Address"
}'
```

```bash
curl -X POST http://localhost:3000/suppliers \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_access_token>" \
-d '{
  "name": "New Supplier",
  "address": "123 Supplier Address"
}'
```

Update Supplier (PUT /suppliers/:id)
```bash
curl -X POST http://localhost:3000/suppliers \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_access_token>" \
-d '{
  "name": "New Supplier",
  "address": "123 Supplier Address"
}'
```

Delete Supplier (DELETE /suppliers/:id)
```bash
curl -X DELETE http://localhost:3000/suppliers/1 \
-H "Authorization: Bearer <your_access_token>"

```
5. Product

A. Get All Products (GET /products)
```bash
curl -X GET http://localhost:3000/products \
-H "Authorization: Bearer <your_access_token>"
```

B. Get Product by ID (GET /products/:id)
```bash
curl -X GET http://localhost:3000/products/1 \
-H "Authorization: Bearer <your_access_token>"
```

C. Create Product (POST /products)
```bash
curl -X POST http://localhost:3000/products \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_access_token>" \
-d '{
  "name": "New Product",
  "price": 1000,
  "stock": 20,
  "category_id": 1,
  "supplier_id": 1
}'
```

D. Update Product (PUT /products/:id)
```bash
curl -X PUT http://localhost:3000/products/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_access_token>" \
-d '{
  "name": "Updated Product",
  "price": 1500,
  "stock": 30,
  "category_id": 2,
  "supplier_id": 2
}'

```
E. Delete Product (DELETE /products/:id)
```bash
curl -X DELETE http://localhost:3000/products/1 \
-H "Authorization: Bearer <your_access_token>"

```
6. Categories

A. Get All Categories (GET /categories)
```bash
curl -X GET http://localhost:3000/categories \
-H "Authorization: Bearer <your_access_token>"

```

B. Get Category by ID (GET /categories/:id)
```bash
curl -X GET http://localhost:3000/categories/1 \
-H "Authorization: Bearer <your_access_token>"

```

C. Create Category (POST /categories)
```bash
curl -X POST http://localhost:3000/categories \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_access_token>" \
-d '{
  "name": "New Category"
}'

```

D. Update Category (PUT /categories/:id)
```bash
curl -X PUT http://localhost:3000/categories/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_access_token>" \
-d '{
  "name": "Updated Category"
}'

```

E. Delete Category (DELETE /categories/:id)
```bash
curl -X DELETE http://localhost:3000/categories/1 \
-H "Authorization: Bearer <your_access_token>"

```
7. Sales

Create Sale (POST /sales)
```bash
curl -X POST http://localhost:3000/sales \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_access_token>" \
-d '{
  "transaction_date": "2024-09-25T14:30:00Z",
  "total_amount": 15000,
  "customer_id": 1,
  "status": "completed",
  "sale_details": [
    {
      "product_id": 1,
      "quantity": 2,
      "price": 5000
    }
  ]
}'

```

8. Payments 

Create Payment (POST /payments)
```bash
curl -X POST http://localhost:3000/payments \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_access_token>" \
-d '{
  "sale_id": 1,
  "amount": 15000,
  "payment_method": "cash",
  "payment_date": "2024-09-25T14:30:00Z"
}'
```