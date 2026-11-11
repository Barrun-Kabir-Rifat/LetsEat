package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	Repo      *repositories.AdminRepository
	JWTSecret []byte
}

type AdminClaims struct {
	AdminID  int64  `json:"admin_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewAdminService(
	repo *repositories.AdminRepository,
	jwtSecret string,
) *AdminService {
	return &AdminService{
		Repo:      repo,
		JWTSecret: []byte(jwtSecret),
	}
}

func (s *AdminService) Register(
	ctx context.Context,
	username string,
	password string,
) (*models.Admin, error) {

	username = strings.TrimSpace(username)

	if username == "" {
		return nil, errors.New("username is required")
	}

	if len(username) > 50 {
		return nil, errors.New("username must not exceed 50 characters")
	}

	if len(password) < 8 {
		return nil, errors.New("password must contain at least 8 characters")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	return s.Repo.Create(ctx, username, string(hash))
}

func (s *AdminService) Login(
	ctx context.Context,
	username string,
	password string,
) (string, error) {

	admin, err := s.Repo.GetByUsername(ctx, strings.TrimSpace(username))

	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(admin.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid username or password")
	}

	claims := AdminClaims{
		AdminID:  admin.ID,
		Username: admin.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "letseat-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(s.JWTSecret)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *AdminService) Logout(
	ctx context.Context,
	jti string,
	expiresAt time.Time,
) error {

	if jti == "" {
		return errors.New("invalid token ID")
	}

	return s.Repo.RevokeToken(ctx, jti, expiresAt)
}
