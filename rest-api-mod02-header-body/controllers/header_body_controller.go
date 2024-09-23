package controllers

import "github.com/gofiber/fiber/v2"

// HeaderBodyController menangani request yang berkaitan dengan header dan body
type HeaderBodyController struct{}

// ResponseFormat adalah struktur untuk format respons standar
type ResponseFormat struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// CreateResourceHandler menangani pembuatan resource baru
// Endpoint: POST /resource
// Headers yang digunakan:
// - Authorization: Token untuk otentikasi
// - Content-Type: tipe konten, misalnya application/json
// - X-Request-ID: ID unik untuk tracking request
func (hbc *HeaderBodyController) CreateResourceHandler(c *fiber.Ctx) error {
	// Mengambil header Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseFormat{
			Code:    "401",
			Message: "Missing Authorization header",
			Data:    nil,
		})
	}

	// Mengambil header Content-Type
	contentType := c.Get("Content-Type")
	if contentType != "application/json" {
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(ResponseFormat{
			Code:    "415",
			Message: "Content-Type must be application/json",
			Data:    nil,
		})
	}

	// Mengambil header X-Request-ID
	requestID := c.Get("X-Request-ID")
	if requestID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFormat{
			Code:    "400",
			Message: "Missing X-Request-ID header",
			Data:    nil,
		})
	}

	// Parse body request
	var resourceData map[string]interface{}
	if err := c.BodyParser(&resourceData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFormat{
			Code:    "400",
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	// Simulasi pembuatan resource (misalnya, simpan ke database)
	createdResource := map[string]interface{}{
		"id":      1, // ID contoh
		"payload": resourceData,
	}

	// Menyiapkan data respons
	responseData := fiber.Map{
		"request_id": requestID,
		"resource":   createdResource,
	}

	// Mengembalikan respons dengan format standar
	return c.Status(fiber.StatusCreated).JSON(ResponseFormat{
		Code:    "201",
		Message: "Resource created successfully",
		Data:    responseData,
	})
}

// GetResourceHandler menangani pengambilan resource berdasarkan ID
// Endpoint: GET /resource/:id
// Headers yang digunakan:
// - Authorization: Token untuk otentikasi
// - Accept: tipe konten yang diterima, misalnya application/json
// - X-Client-Version: versi client yang mengirim request
func (hbc *HeaderBodyController) GetResourceHandler(c *fiber.Ctx) error {
	// Mengambil header Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseFormat{
			Code:    "401",
			Message: "Missing Authorization header",
			Data:    nil,
		})
	}

	// Mengambil header Accept
	acceptHeader := c.Get("Accept")
	if acceptHeader != "application/json" {
		return c.Status(fiber.StatusNotAcceptable).JSON(ResponseFormat{
			Code:    "406",
			Message: "Accept header must be application/json",
			Data:    nil,
		})
	}

	// Mengambil header X-Client-Version
	clientVersion := c.Get("X-Client-Version")
	if clientVersion == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFormat{
			Code:    "400",
			Message: "Missing X-Client-Version header",
			Data:    nil,
		})
	}

	// Mengambil parameter ID dari URL
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFormat{
			Code:    "400",
			Message: "Resource ID is required",
			Data:    nil,
		})
	}

	// Simulasi pengambilan resource (misalnya, dari database)
	resource := map[string]interface{}{
		"id":   id,
		"name": "Sample Resource",
	}

	// Menyiapkan data respons
	responseData := fiber.Map{
		"client_version": clientVersion,
		"resource":       resource,
	}

	// Mengembalikan respons dengan format standar
	return c.JSON(ResponseFormat{
		Code:    "200",
		Message: "Resource retrieved successfully",
		Data:    responseData,
	})
}

// UpdateResourceHandler menangani pembaruan resource berdasarkan ID
// Endpoint: PUT /resource/:id
// Headers yang digunakan:
// - Authorization: Token untuk otentikasi
// - Content-Type: tipe konten, misalnya application/json
// - X-Request-ID: ID unik untuk tracking request
func (hbc *HeaderBodyController) UpdateResourceHandler(c *fiber.Ctx) error {
	// Mengambil header Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseFormat{
			Code:    "401",
			Message: "Missing Authorization header",
			Data:    nil,
		})
	}

	// Mengambil header Content-Type
	contentType := c.Get("Content-Type")
	if contentType != "application/json" {
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(ResponseFormat{
			Code:    "415",
			Message: "Content-Type must be application/json",
			Data:    nil,
		})
	}

	// Mengambil header X-Request-ID
	requestID := c.Get("X-Request-ID")
	if requestID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFormat{
			Code:    "400",
			Message: "Missing X-Request-ID header",
			Data:    nil,
		})
	}

	// Mengambil parameter ID dari URL
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFormat{
			Code:    "400",
			Message: "Resource ID is required",
			Data:    nil,
		})
	}

	// Parse body request
	var updateData map[string]interface{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFormat{
			Code:    "400",
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	// Simulasi pembaruan resource (misalnya, update di database)
	updatedResource := map[string]interface{}{
		"id":      id,
		"payload": updateData,
	}

	// Menyiapkan data respons
	responseData := fiber.Map{
		"request_id": requestID,
		"resource":   updatedResource,
	}

	// Mengembalikan respons dengan format standar
	return c.JSON(ResponseFormat{
		Code:    "200",
		Message: "Resource updated successfully",
		Data:    responseData,
	})
}

// DeleteResourceHandler menangani penghapusan resource berdasarkan ID
// Endpoint: DELETE /resource/:id
// Headers yang digunakan:
// - Authorization: Token untuk otentikasi
// - X-Request-ID: ID unik untuk tracking request
// - X-Client-Version: versi client yang mengirim request
func (hbc *HeaderBodyController) DeleteResourceHandler(c *fiber.Ctx) error {
	// Mengambil header Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseFormat{
			Code:    "401",
			Message: "Missing Authorization header",
			Data:    nil,
		})
	}

	// Mengambil header X-Request-ID
	requestID := c.Get("X-Request-ID")
	if requestID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFormat{
			Code:    "400",
			Message: "Missing X-Request-ID header",
			Data:    nil,
		})
	}

	// Mengambil header X-Client-Version
	clientVersion := c.Get("X-Client-Version")
	if clientVersion == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFormat{
			Code:    "400",
			Message: "Missing X-Client-Version header",
			Data:    nil,
		})
	}

	// Mengambil parameter ID dari URL
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFormat{
			Code:    "400",
			Message: "Resource ID is required",
			Data:    nil,
		})
	}

	// Simulasi penghapusan resource (misalnya, hapus dari database)
	deleted := true // Asumsi penghapusan berhasil

	if !deleted {
		return c.Status(fiber.StatusNotFound).JSON(ResponseFormat{
			Code:    "404",
			Message: "Resource not found",
			Data:    nil,
		})
	}

	// Menyiapkan data respons
	responseData := fiber.Map{
		"request_id":     requestID,
		"client_version": clientVersion,
		"deleted_id":     id,
	}

	// Mengembalikan respons dengan format standar
	return c.JSON(ResponseFormat{
		Code:    "200",
		Message: "Resource deleted successfully",
		Data:    responseData,
	})
}

// SearchResourceHandler menangani pencarian resource berdasarkan query
// Endpoint: GET /resource/search
// Headers yang digunakan:
// - Authorization: Token untuk otentikasi
// - Accept: tipe konten yang diterima, misalnya application/json
// - X-Search-Params: parameter pencarian tambahan
func (hbc *HeaderBodyController) SearchResourceHandler(c *fiber.Ctx) error {
	// Mengambil header Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseFormat{
			Code:    "401",
			Message: "Missing Authorization header",
			Data:    nil,
		})
	}

	// Mengambil header Accept
	acceptHeader := c.Get("Accept")
	if acceptHeader != "application/json" {
		return c.Status(fiber.StatusNotAcceptable).JSON(ResponseFormat{
			Code:    "406",
			Message: "Accept header must be application/json",
			Data:    nil,
		})
	}

	// Mengambil header X-Search-Params
	searchParams := c.Get("X-Search-Params")
	if searchParams == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFormat{
			Code:    "400",
			Message: "Missing X-Search-Params header",
			Data:    nil,
		})
	}

	// Mengambil query parameter
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseFormat{
			Code:    "400",
			Message: "Missing query parameter 'q'",
			Data:    nil,
		})
	}

	// Simulasi pencarian resource (misalnya, cari di database)
	results := []map[string]interface{}{
		{
			"id":   1,
			"name": "Sample Resource 1",
		},
		{
			"id":   2,
			"name": "Sample Resource 2",
		},
	}

	// Menyiapkan data respons
	responseData := fiber.Map{
		"search_params": searchParams,
		"query":         query,
		"results":       results,
	}

	// Mengembalikan respons dengan format standar
	return c.JSON(ResponseFormat{
		Code:    "200",
		Message: "Search completed successfully",
		Data:    responseData,
	})
}
