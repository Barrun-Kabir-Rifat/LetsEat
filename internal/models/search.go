package models

type MenuItemSearchResult struct {
	ID             int64   `json:"id"`
	RestaurantID   int64   `json:"restaurant_id"`
	RestaurantName string  `json:"restaurant_name"`
	CategoryID     int64   `json:"category_id"`
	CategoryName   string  `json:"category_name"`
	Name           string  `json:"name"`
	Description    *string `json:"description,omitempty"`
	Price          string  `json:"price"`
	IsAvailable    bool    `json:"is_available"`
}
