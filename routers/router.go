package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nazordz/boutique-product-catalog-service/routers/handlers"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/catalogs", handlers.GetCatalogs)
	r.GET("/catalogs/:id", handlers.GetCatalogById)
	r.POST("/catalogs", handlers.CreateCatalog)

	r.GET("/products", handlers.GetProducts)
	r.POST("/products", handlers.CreateProduct)

	return r
}
