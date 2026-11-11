package repositories

import (
	"context"
	"time"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminRepository struct {
	DB *pgxpool.Pool
}

func NewAdminRepository(db *pgxpool.Pool) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (r *AdminRepository) Create(
	ctx context.Context,
	username string,
	passwordHash string,
) (*models.Admin, error) {

	admin := &models.Admin{}

	err := r.DB.QueryRow(
		ctx,
		`INSERT INTO admins (username, password_hash)
         VALUES ($1, $2)
         RETURNING id, username, password_hash, created_at, updated_at`,
		username,
		passwordHash,
	).Scan(
		&admin.ID,
		&admin.Username,
		&admin.PasswordHash,
		&admin.CreatedAt,
		&admin.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *AdminRepository) GetByUsername(
	ctx context.Context,
	username string,
) (*models.Admin, error) {

	admin := &models.Admin{}

	err := r.DB.QueryRow(
		ctx,
		`SELECT id, username, password_hash, created_at, updated_at
         FROM admins
         WHERE LOWER(username) = LOWER($1)`,
		username,
	).Scan(
		&admin.ID,
		&admin.Username,
		&admin.PasswordHash,
		&admin.CreatedAt,
		&admin.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *AdminRepository) RevokeToken(
	ctx context.Context,
	jti string,
	expiresAt time.Time,
) error {

	_, err := r.DB.Exec(
		ctx,
		`INSERT INTO revoked_tokens (jti, expires_at)
         VALUES ($1, $2)
         ON CONFLICT (jti) DO NOTHING`,
		jti,
		expiresAt,
	)

	return err
}

func (r *AdminRepository) IsTokenRevoked(
	ctx context.Context,
	jti string,
) (bool, error) {

	var revoked bool

	err := r.DB.QueryRow(
		ctx,
		`SELECT EXISTS (
            SELECT 1
            FROM revoked_tokens
            WHERE jti = $1
              AND expires_at > NOW()
        )`,
		jti,
	).Scan(&revoked)

	return revoked, err
}
