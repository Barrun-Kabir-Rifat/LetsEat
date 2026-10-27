package repositories

import (
	"context"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RestaurantRepository struct {
	DB *pgxpool.Pool
}

func NewRestaurantRepository(db *pgxpool.Pool) *RestaurantRepository {
	return &RestaurantRepository{DB: db}
}

func (r *RestaurantRepository) Create(
	ctx context.Context,
	restaurant *models.Restaurant,
) error {
	query := `
		INSERT INTO restaurants (
			name, address, phone, description, is_active
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	return r.DB.QueryRow(
		ctx,
		query,
		restaurant.Name,
		restaurant.Address,
		restaurant.Phone,
		restaurant.Description,
		restaurant.IsActive,
	).Scan(
		&restaurant.ID,
		&restaurant.CreatedAt,
		&restaurant.UpdatedAt,
	)
}

func (r *RestaurantRepository) GetAll(
	ctx context.Context,
) ([]models.Restaurant, error) {
	query := `
        SELECT id, name, address, phone, description,
               is_active, created_at, updated_at
        FROM restaurants
        ORDER BY id
    `

	rows, err := r.DB.Query(ctx, query)
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
