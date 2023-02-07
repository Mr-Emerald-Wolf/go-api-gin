package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string //`json:"name"`
	Price       int    //`json:"price"`
	Seller      string //`json:"seller"`
	Description string //`json:"desc"`
}

type NewProduct struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Seller      string `json:"seller"`
	Description string `json:"desc"`
}