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

type FlashSaleProduct struct {
	FlashSaleId uint32 `json:"flash_sale_id" gorm:"primaryKey"` // Primary Key

	ProductID          uint32    `json:"product_id"`
	ProductName        string    `json:"product_name"`
	ProductDescription string    `json:"product_description"`
	Price              float64   `json:"price"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`

	Discount  float64   `json:"discount"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (f *FlashSaleProduct) TableName() string {
	return "flash_sale"
}
