package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nazordz/boutique-product-catalog-service/dtos"
	"github.com/nazordz/boutique-product-catalog-service/models"
)

func GetProducts(c *gin.Context) {
	products, err := models.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var form dtos.CreateProductRequest
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}
	newProduct := models.Product{
		CatalogID:     form.CatalogId,
		Name:          form.Name,
		Description:   form.Description,
		Price:         form.Price,
		StockQuantity: form.StockQuantity,
	}
	err := models.CreateProduct(&newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, newProduct)
}
