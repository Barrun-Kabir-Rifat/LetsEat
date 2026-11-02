package repositories

import (
	"context"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type MenuItemRepository struct {
	DB *pgxpool.Pool
}

func NewMenuItemRepository(db *pgxpool.Pool) *MenuItemRepository {
	return &MenuItemRepository{DB: db}
}

func (r *MenuItemRepository) Create(
	ctx context.Context,
	menuItem *models.MenuItem,
) error {
	query := `
        INSERT INTO menu_items (
            restaurant_id,
            category_id,
            name,
            description,
            price,
            is_available
        )
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at, updated_at
    `

	return r.DB.QueryRow(
		ctx,
		query,
		menuItem.RestaurantID,
		menuItem.CategoryID,
		menuItem.Name,
		menuItem.Description,
		menuItem.Price,
		menuItem.IsAvailable,
	).Scan(
		&menuItem.ID,
		&menuItem.CreatedAt,
		&menuItem.UpdatedAt,
	)
}

func (r *MenuItemRepository) GetAllByRestaurantID(
	ctx context.Context,
	restaurantID int64,
) ([]models.MenuItem, error) {
	query := `
        SELECT
            id,
            restaurant_id,
            category_id,
            name,
            description,
            price::TEXT,
            is_available,
            created_at,
            updated_at
        FROM menu_items
        WHERE restaurant_id = $1
        ORDER BY id
    `

	rows, err := r.DB.Query(ctx, query, restaurantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	menuItems := make([]models.MenuItem, 0)

	for rows.Next() {
		var item models.MenuItem
		var priceText string

		err := rows.Scan(
			&item.ID,
			&item.RestaurantID,
			&item.CategoryID,
			&item.Name,
			&item.Description,
			&priceText,
			&item.IsAvailable,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		item.Price, err = decimal.NewFromString(priceText)
		if err != nil {
			return nil, err
		}

		menuItems = append(menuItems, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return menuItems, nil
}

func (r *MenuItemRepository) GetByID(
	ctx context.Context,
	id int64,
) (*models.MenuItem, error) {
	query := `
        SELECT id, restaurant_id, category_id,
               name, description, price::TEXT,
               is_available, created_at, updated_at
        FROM menu_items
        WHERE id = $1
    `

	var item models.MenuItem
	var priceText string

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&item.ID,
		&item.RestaurantID,
		&item.CategoryID,
		&item.Name,
		&item.Description,
		&priceText,
		&item.IsAvailable,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	item.Price, err = decimal.NewFromString(priceText)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *MenuItemRepository) Update(
	ctx context.Context,
	item *models.MenuItem,
) error {
	query := `
        UPDATE menu_items
        SET category_id = $1,
            name = $2,
            description = $3,
            price = $4::numeric,
            is_available = $5,
            updated_at = NOW()
        WHERE id = $6
        RETURNING restaurant_id, created_at, updated_at
    `

	return r.DB.QueryRow(
		ctx,
		query,
		item.CategoryID,
		item.Name,
		item.Description,
		item.Price.String(),
		item.IsAvailable,
		item.ID,
	).Scan(
		&item.RestaurantID,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
}

func (r *MenuItemRepository) UpdatePrice(
	ctx context.Context,
	id int64,
	price decimal.Decimal,
) error {
	query := `
        UPDATE menu_items
        SET price = $1::numeric,
            updated_at = NOW()
        WHERE id = $2
        RETURNING id
    `

	var updatedID int64

	return r.DB.QueryRow(
		ctx, query, price.String(), id,
	).Scan(&updatedID)
}

func (r *MenuItemRepository) UpdateAvailability(
	ctx context.Context,
	id int64,
	isAvailable bool,
) error {
	query := `
        UPDATE menu_items
        SET is_available = $1,
            updated_at = NOW()
        WHERE id = $2
        RETURNING id
    `

	var updatedID int64

	return r.DB.QueryRow(
		ctx, query, isAvailable, id,
	).Scan(&updatedID)
}

func (r *MenuItemRepository) Delete(
	ctx context.Context,
	id int64,
) error {
	query := `
        DELETE FROM menu_items
        WHERE id = $1
        RETURNING id
    `

	var deletedID int64

	return r.DB.QueryRow(
		ctx, query, id,
	).Scan(&deletedID)
}
