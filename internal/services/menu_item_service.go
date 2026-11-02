package services

import (
	"context"
	"errors"
	"strings"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/repositories"
	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

var ErrMenuItemNotFound = errors.New("food item not found")

type MenuItemService struct {
	Repository *repositories.MenuItemRepository
}

func NewMenuItemService(
	repository *repositories.MenuItemRepository,
) *MenuItemService {
	return &MenuItemService{
		Repository: repository,
	}
}

func (s *MenuItemService) Create(
	ctx context.Context,
	menuItem *models.MenuItem,
) error {
	menuItem.Name = strings.TrimSpace(menuItem.Name)

	if menuItem.RestaurantID <= 0 {
		return errors.New("invalid restaurant ID")
	}

	if menuItem.CategoryID <= 0 {
		return errors.New("invalid category ID")
	}

	if menuItem.Name == "" {
		return errors.New("food item name is required")
	}

	if !menuItem.Price.IsPositive() {
		return errors.New("food item price must be greater than zero")
	}

	return s.Repository.Create(ctx, menuItem)
}

func (s *MenuItemService) GetAllByRestaurantID(
	ctx context.Context,
	restaurantID int64,
) ([]models.MenuItem, error) {
	if restaurantID <= 0 {
		return nil, errors.New("invalid restaurant ID")
	}

	return s.Repository.GetAllByRestaurantID(ctx, restaurantID)
}

func (s *MenuItemService) GetByID(
	ctx context.Context,
	id int64,
) (*models.MenuItem, error) {
	if id <= 0 {
		return nil, errors.New("invalid food item ID")
	}

	item, err := s.Repository.GetByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrMenuItemNotFound
	}

	return item, err
}

func (s *MenuItemService) Update(
	ctx context.Context,
	item *models.MenuItem,
) error {
	if item.ID <= 0 {
		return errors.New("invalid food item ID")
	}

	item.Name = strings.TrimSpace(item.Name)

	if item.Name == "" {
		return errors.New("food item name is required")
	}

	if item.CategoryID <= 0 {
		return errors.New("invalid category ID")
	}

	if !item.Price.IsPositive() {
		return errors.New("food item price must be greater than zero")
	}

	err := s.Repository.Update(ctx, item)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrMenuItemNotFound
	}

	return err
}

func (s *MenuItemService) UpdatePrice(
	ctx context.Context,
	id int64,
	price decimal.Decimal,
) error {
	if id <= 0 {
		return errors.New("invalid food item ID")
	}

	if !price.IsPositive() {
		return errors.New("food item price must be greater than zero")
	}

	err := s.Repository.UpdatePrice(ctx, id, price)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrMenuItemNotFound
	}

	return err
}

func (s *MenuItemService) UpdateAvailability(
	ctx context.Context,
	id int64,
	isAvailable bool,
) error {
	if id <= 0 {
		return errors.New("invalid food item ID")
	}

	err := s.Repository.UpdateAvailability(
		ctx, id, isAvailable,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrMenuItemNotFound
	}

	return err
}

func (s *MenuItemService) Delete(
	ctx context.Context,
	id int64,
) error {
	if id <= 0 {
		return errors.New("invalid food item ID")
	}

	err := s.Repository.Delete(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrMenuItemNotFound
	}

	return err
}
