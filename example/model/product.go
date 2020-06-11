package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Sku string
	Price int
}