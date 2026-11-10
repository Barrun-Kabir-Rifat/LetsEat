package repositories

import (
    "context"

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