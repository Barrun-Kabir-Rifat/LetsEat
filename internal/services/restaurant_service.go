package services

import (
	"context"
	"errors"
	"strings"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/repositories"
)

type RestaurantService struct {
	Repository *repositories.RestaurantRepository
}

func NewRestaurantService(
	repository *repositories.RestaurantRepository,
) *RestaurantService {
	return &RestaurantService{Repository: repository}
}

func (s *RestaurantService) Create(
	ctx context.Context,
	restaurant *models.Restaurant,
) error {
	restaurant.Name = strings.TrimSpace(restaurant.Name)
	restaurant.Address = strings.TrimSpace(restaurant.Address)

	if restaurant.Name == "" {
		return errors.New("restaurant name is required")
	}

	if restaurant.Address == "" {
		return errors.New("restaurant address is required")
	}

	return s.Repository.Create(ctx, restaurant)
}

func (s *RestaurantService) GetAll(
	ctx context.Context,
) ([]models.Restaurant, error) {
	return s.Repository.GetAll(ctx)
}
