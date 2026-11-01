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
