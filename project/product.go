package project

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	SerialNumber string  `json:"sn"`
	Name         string  `json:"name"`
	Description  string  `json:"descrition"`
	Price        float64 `json:"price"`
}

var products = []Product{
	{
		SerialNumber: "SN1001",
		Name:         "Keyboard",
		Description:  "Mechanical keyboard",
		Price:        99.9,
	},
	{
		SerialNumber: "SN1002",
		Name:         "Mouse",
		Description:  "Wireless mouse",
		Price:        49.5,
	},
	{
		SerialNumber: "SN1003",
		Name:         "Monitor",
		Description:  "27 inch 4K",
		Price:        399.0,
	},
}

// curl to test product endpoints:
//
//  1. Create a Product:
//     curl -X POST http://localhost:8080/api/v1/products \
//     -H "Content-Type: application/json" \
//     -d '{"sn":"SN2001", "name":"Laptop", "description":"13 inch", "price":999.9}'
//
//  2. List all products: curl http://localhost:8080/api/v1/products
//
//  3. Update a product:
//     curl -X PATCH http://localhost:8080/api/v1/products/SN2001 \
//     -H "Content-Type: application/json" \
//     -d '{"price":888.8}'
//
//  4. Delete a product: curl -X DELETE http://localhost:8080/api/v1/products/SN2001
func ProductAPITest() {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/products", getProducts)
		api.GET("/products/:sn", getProduct)
		api.POST("/products", createProduct)
		api.PATCH("/products/:sn", updateProduct)
		api.DELETE("/products/:sn", deleteProduct)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "route not found",
			"path":  c.Request.URL.Path,
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server, %v", err)
	}
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {
	sn := c.Param("sn")

	for _, p := range products {
		if p.SerialNumber == sn {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
}

type CreateProductInput struct {
	SerialNumber string  `json:"sn" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Description  string  `json:"descrition"`
	Price        float64 `json:"price" binding:"required,gte=0"`
}

func createProduct(c *gin.Context) {
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := Product{
		SerialNumber: input.SerialNumber,
		Name:         input.Name,
		Description:  input.Description,
		Price:        input.Price,
	}
	products = append(products, p)

	c.JSON(http.StatusCreated, p)
}

type UpdateProductInput struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price" binding:"omitempty,gte=0"`
}

func updateProduct(c *gin.Context) {
	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	sn := c.Param("sn")

	for i, p := range products {
		if p.SerialNumber == sn {
			// only update if provided
			if input.Name != nil {
				products[i].Name = *input.Name
			}
			if input.Description != nil {
				products[i].Description = *input.Description
			}
			if input.Price != nil {
				products[i].Price = *input.Price
			}

			c.JSON(http.StatusOK, products[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
}

func deleteProduct(c *gin.Context) {
	sn := c.Param("sn")

	for i, p := range products {
		if p.SerialNumber == sn {
			// remove element from slice
			products = append(products[:i], products[i+1:]...)

			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
}
