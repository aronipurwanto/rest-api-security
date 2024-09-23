// gateway/gateway.go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"log"
)

func main() {
	app := fiber.New()

	// Proxy to Product Service
	app.Get("/products", func(c *fiber.Ctx) error {
		// Proxy request to service 1 (Product Service)
		return proxy.Do(c, "http://localhost:3001/products")
	})

	// Proxy to Order Service
	app.Get("/orders", func(c *fiber.Ctx) error {
		// Proxy request to service 2 (Order Service)
		return proxy.Do(c, "http://localhost:3002/orders")
	})

	// Proxy to Customer Service
	app.Get("/customers", func(c *fiber.Ctx) error {
		// Proxy request to service 3 (Customer Service)
		return proxy.Do(c, "http://localhost:3003/customers")
	})

	log.Fatal(app.Listen(":3000"))
}
