package services

import (
	"context"
	"errors"
	"strings"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/repositories"
	"github.com/jackc/pgx/v5"
)

type CategoryService struct {
	Repository *repositories.CategoryRepository
}

func NewCategoryService(
	repository *repositories.CategoryRepository,
) *CategoryService {
	return &CategoryService{Repository: repository}
}

func (s *CategoryService) Create(
	ctx context.Context,
	category *models.Category,
) error {
	category.Name = strings.TrimSpace(category.Name)

	if category.Name == "" {
		return errors.New("category name is required")
	}

	return s.Repository.Create(ctx, category)
}

func (s *CategoryService) GetAll(
	ctx context.Context,
) ([]models.Category, error) {
	return s.Repository.GetAll(ctx)
}

var ErrCategoryNotFound = errors.New("category not found")

func (s *CategoryService) GetByID(
	ctx context.Context,
	id int64,
) (*models.Category, error) {
	category, err := s.Repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrCategoryNotFound
		}

		return nil, err
	}

	return category, nil
}
