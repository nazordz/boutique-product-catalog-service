package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nazordz/boutique/product-catalog-service/dtos"
	"github.com/nazordz/boutique/product-catalog-service/models"
)

func GetCatalogs(c *gin.Context) {
	catalogs, err := models.GetCatalogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": "Error",
		})
		return
	}

	c.JSON(http.StatusOK, catalogs)
}

func CreateCatalog(c *gin.Context) {
	var form dtos.CreateCatalogRequest
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messages": err.Error()})
		return
	}

	catalog := models.Catalog{
		Name:        form.Name,
		Description: form.Description,
	}

	err := models.CreateCatalog(&catalog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, catalog)
}

func GetCatalogById(c *gin.Context) {
	var uri dtos.CatalogUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}
	catalog, err := models.GetCatalog(uri.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, catalog)
}
