package services

import (
	"context"
	"errors"
	"strings"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/repositories"
	"github.com/jackc/pgx/v5"
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

var ErrRestaurantNotFound = errors.New("restaurant not found")

func (s *RestaurantService) GetByID(
	ctx context.Context,
	id int64,
) (*models.Restaurant, error) {
	restaurant, err := s.Repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrRestaurantNotFound
		}

		return nil, err
	}

	return restaurant, nil
}

func (s *RestaurantService) Update(
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

	if restaurant.ID <= 0 {
		return errors.New("invalid restaurant ID")
	}

	err := s.Repository.Update(ctx, restaurant)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrRestaurantNotFound
		}

		return err
	}

	return nil

}

func (s *RestaurantService) Delete(
	ctx context.Context,
	id int64,
) error {
	if id <= 0 {
		return errors.New("invalid restaurant ID")
	}

	err := s.Repository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrRestaurantNotFound
		}

		return err
	}

	return nil
}
