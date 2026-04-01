package models

import (
	"time"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

var (
	OrderStatuses = []string{"Order Placed", "Preparing", "Baking", "Quality Check", "Ready"}

	PizzaTypes = []string{
		"Margherita",
		"Pepperoni",
		"BBQ Chicken",
		"Veggie",
		"Hawaiian",
		"Buffalo Chicken",
		"Supreme",
		"Truffle Mushroom",
		"Four Cheese",
	}

	PizzaSizes = []string{"Small", "Medium", "Large", "X-Large"}
)

type OrderModels struct {
	DB *gorm.DB
}

type Order struct {
	ID           string      `gorm:"primaryKey;size:14" json:"id"`
	Status       string      `gorm:"not null" json:"status"`
	CustomerName string      `gorm:"not null" json:"customerName"`
	Phone        string      `gorm:"not null" json:"phone"`
	Address      string      `gorm:"not null" json:"address"`
	Items        []OrderItem `gorm:"foreignKey:OrderID" json:"pizzas"`
	CreatedAt    time.Time   `json:"createdAt"`
}

type OrderItem struct {
	Id           string `gorm:"primaryKey;size:14" json:"id"`
	OrderID      string `gorm:"index;size:14" json:"orderId"`
	Size         string `gorm:"not null" json:"size"`
	Pizza        string `gorm:"not null" json:"pizza"`
	Instructions string `json:"instructions"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.ID == "" {
		o.ID = shortid.MustGenerate()
	}
	return nil
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if oi.Id == "" {
		oi.Id = shortid.MustGenerate()
	}
	return nil

}

func (o *OrderModels) CreateOrder(order *Order) error {
	return o.DB.Create(order).Error
}

func (o *OrderModels) GetOrder(id string) (*Order, error) {
	var order Order
	err := o.DB.Preload("Items").First(&order, "id = ?", id).Error
	return &order, err
}
