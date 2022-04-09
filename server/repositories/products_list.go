package repository

import (
	"github.com/lil-shimon/oms/db"
	"github.com/lil-shimon/oms/models"
)

// products lists type
type ProductsLists []*models.ProductsList

// GetProductslists is a function to get products lists
func GetProductsLists() (ProductsLists, error) {
	i := ProductsLists{}

	if err := db.DB.Find(&i).Error; err != nil {
		return i, err
	}

	return i, nil
}
