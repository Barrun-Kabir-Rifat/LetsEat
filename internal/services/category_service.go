package services

import (
	"context"
	"errors"
	"strings"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/repositories"
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
