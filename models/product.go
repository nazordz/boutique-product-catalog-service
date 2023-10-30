package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID            string    `gorm:"primaryKey" json:"id"`
	CatalogID     string    `json:"catalog_id"`
	Catalog       Catalog   `json:"catalog"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         float32   `json:"price"`
	StockQuantity int       `json:"stock_quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) error {
	product.ID = uuid.NewString()
	return nil
}

func GetProducts() ([]*Product, error) {
	var products []*Product
	err := db.Preload("Catalog").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func CreateProduct(product *Product) error {
	err := db.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}
