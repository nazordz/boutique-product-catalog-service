package dtos

type CreateCatalogRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateProductRequest struct {
	CatalogId     string  `json:"catalog_id" binding:"required"`
	Name          string  `json:"name" binding:"required"`
	Description   string  `json:"description"`
	Price         float32 `json:"price" binding:"required"`
	StockQuantity int     `json:"stock_quantity" binding:"required"`
}
