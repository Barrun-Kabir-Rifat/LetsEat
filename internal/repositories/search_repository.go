package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SearchRepository struct {
	DB *pgxpool.Pool
}

func NewSearchRepository(db *pgxpool.Pool) *SearchRepository {
	return &SearchRepository{DB: db}
}

// SearchRestaurants searches restaurants by name.
func (r *SearchRepository) SearchRestaurants(
	ctx context.Context,
	name string,
) ([]models.Restaurant, error) {
	query := `
        SELECT id, name, address, phone, description,
               is_active, created_at, updated_at
        FROM restaurants
        WHERE name ILIKE $1
          AND is_active = TRUE
        ORDER BY name
    `

	rows, err := r.DB.Query(ctx, query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	restaurants := make([]models.Restaurant, 0)

	for rows.Next() {
		var restaurant models.Restaurant

		err := rows.Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Address,
			&restaurant.Phone,
			&restaurant.Description,
			&restaurant.IsActive,
			&restaurant.CreatedAt,
			&restaurant.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		restaurants = append(restaurants, restaurant)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return restaurants, nil
}

// SearchMenuItems searches food items using optional filters.
func (r *SearchRepository) SearchMenuItems(
	ctx context.Context,
	name string,
	categoryID int64,
	restaurantID int64,
	available bool,
) ([]models.MenuItemSearchResult, error) {

	query := `
        SELECT
            mi.id,
            mi.restaurant_id,
            r.name AS restaurant_name,
            mi.category_id,
            c.name AS category_name,
            mi.name,
            mi.description,
            mi.price::TEXT,
            mi.is_available
        FROM menu_items mi
        INNER JOIN restaurants r
            ON r.id = mi.restaurant_id
        INNER JOIN categories c
            ON c.id = mi.category_id
        WHERE r.is_active = TRUE
          AND mi.is_available = $1
    `

	args := []interface{}{available}
	var conditions strings.Builder

	if name != "" {
		args = append(args, "%"+name+"%")
		fmt.Fprintf(
			&conditions,
			" AND mi.name ILIKE $%d",
			len(args),
		)
	}

	if categoryID > 0 {
		args = append(args, categoryID)
		fmt.Fprintf(
			&conditions,
			" AND mi.category_id = $%d",
			len(args),
		)
	}

	if restaurantID > 0 {
		args = append(args, restaurantID)
		fmt.Fprintf(
			&conditions,
			" AND mi.restaurant_id = $%d",
			len(args),
		)
	}

	query += conditions.String()
	query += " ORDER BY r.name, mi.name"

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]models.MenuItemSearchResult, 0)

	for rows.Next() {
		var item models.MenuItemSearchResult

		err := rows.Scan(
			&item.ID,
			&item.RestaurantID,
			&item.RestaurantName,
			&item.CategoryID,
			&item.CategoryName,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.IsAvailable,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
