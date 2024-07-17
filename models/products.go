package model

import (
	"github.com/google/uuid"
)

type Products struct {
	ID         uuid.UUID `gorm:"type:uuid; primaryKey"`
	Name       string    `json:"name" validate:"required"`
	Desc       string    `json:"desc" validate:"required"`
	Price      float64   `json:"price" validate:"required"`
	CategoryID uuid.UUID `json:"category_id" validate:"required"`
	Model
}

type Categories struct {
	ID   uuid.UUID `gorm:"type:uuid; primaryKey"`
	Name string    `json:"name"`
	Model
}

type CartItems struct {
	ID        uuid.UUID `gorm:"type:uuid; primaryKey"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	Quantity  uuid.UUID `json:"quantity" validate:"required"`
	Model
}

type Orders struct {
	ID             uuid.UUID `gorm:"type:uuid; primaryKey"`
	UserID         uuid.UUID `json:"user_id" validate:"required"`
	TotalAmount    float64   `json:"total_amount" validate:"required"`
	Status         string    `json:"status" validate:"required"`
	PaymentMethod  string    `json:"payment_method" validate:"required"`
	BillingAddress string    `json:"billing_address" validate:"required"`
	Model
}

type OrderItems struct {
	ID           uuid.UUID `gorm:"type:uuid; primaryKey"`
	OrderID      uuid.UUID `json:"order_id" validate:"required"`
	ProductID    uuid.UUID `json:"product_id" validate:"required"`
	ProductName  string    `json:"product_name" validate:"required"`
	ProductPrice float64   `json:"product_price" validate:"required"`
	Quantity     uuid.UUID `json:"quantity" validate:"required"`
	SubTotal     float64   `json:"sub_total"`
	Model
}
