package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Catalog struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"products"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (catalog *Catalog) BeforeCreate(tx *gorm.DB) error {
	catalog.ID = uuid.NewString()
	return nil
}

func GetCatalogs() ([]*Catalog, error) {
	var catalogs []*Catalog
	err := db.Find(&catalogs).Error
	if err != nil {
		return nil, err
	}
	return catalogs, err
}

func GetCatalog(id string) (*Catalog, error) {
	var catalog *Catalog

	err := db.Preload("Products").First(&catalog, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return catalog, nil
}

func CreateCatalog(catalog *Catalog) error {
	err := db.Create(&catalog).Error
	if err != nil {
		return err
	}
	return nil
}
