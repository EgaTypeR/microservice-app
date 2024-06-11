package models

import "time"

type Product struct {
	ProductID          uint32    `json:"product_id" gorm:"primaryKey"`
	ProductName        string    `json:"product_name"`
	ProductDescription string    `json:"product_description"`
	Price              float64   `json:"price"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (p *Product) TableName() string {
	return "product"
}
