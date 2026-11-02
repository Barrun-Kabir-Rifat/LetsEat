package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type MenuItem struct {
	ID           int64           `json:"id"`
	RestaurantID int64           `json:"restaurant_id"`
	CategoryID   int64           `json:"category_id"`
	Name         string          `json:"name"`
	Description  *string         `json:"description,omitempty"`
	Price        decimal.Decimal `json:"price"`
	IsAvailable  bool            `json:"is_available"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}
