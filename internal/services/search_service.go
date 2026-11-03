package services

import (
	"context"
	"errors"
	"strings"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/repositories"
)

type SearchService struct {
	Repository *repositories.SearchRepository
}

func NewSearchService(
	repository *repositories.SearchRepository,
) *SearchService {
	return &SearchService{
		Repository: repository,
	}
}

// SearchRestaurants searches restaurants by name.
func (s *SearchService) SearchRestaurants(
	ctx context.Context,
	name string,
) ([]models.Restaurant, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, errors.New("restaurant search name is required")
	}

	return s.Repository.SearchRestaurants(ctx, name)
}

// SearchMenuItems searches food items using optional filters.
func (s *SearchService) SearchMenuItems(
	ctx context.Context,
	name string,
	categoryID int64,
	restaurantID int64,
	available bool,
) ([]models.MenuItemSearchResult, error) {
	name = strings.TrimSpace(name)

	if categoryID < 0 {
		return nil, errors.New("invalid category ID")
	}

	if restaurantID < 0 {
		return nil, errors.New("invalid restaurant ID")
	}

	return s.Repository.SearchMenuItems(
		ctx,
		name,
		categoryID,
		restaurantID,
		available,
	)
}
