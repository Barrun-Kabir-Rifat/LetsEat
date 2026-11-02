package repositories

import (
	"context"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository struct {
	DB *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) Create(
	ctx context.Context,
	category *models.Category,
) error {
	query := `
		INSERT INTO categories (name)
		VALUES ($1)
		RETURNING id, created_at, updated_at
	`

	return r.DB.QueryRow(
		ctx,
		query,
		category.Name,
	).Scan(
		&category.ID,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
}

func (r *CategoryRepository) GetAll(
	ctx context.Context,
) ([]models.Category, error) {
	query := `
        SELECT id, name, created_at, updated_at
        FROM categories
        ORDER BY id
    `

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]models.Category, 0)

	for rows.Next() {
		var category models.Category

		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) GetByID(
	ctx context.Context,
	id int64,
) (*models.Category, error) {
	query := `
        SELECT id, name, created_at, updated_at
        FROM categories
        WHERE id = $1
    `

	var category models.Category

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &category, nil
}
