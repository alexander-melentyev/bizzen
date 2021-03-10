package repository

import (
	"context"

	"github.com/alexander-melentyev/bizzen/internal/domain"
	"github.com/jmoiron/sqlx"
)

// Repository -.
type Repository struct {
	conn *sqlx.DB
}

var _ domain.OrgRepository = (*Repository)(nil)

// NewRepository -.
func NewRepository(conn *sqlx.DB) *Repository {
	return &Repository{
		conn: conn,
	}
}

// Create -.
func (r *Repository) Create(ctx context.Context, o domain.Org) error {
	return nil
}

// ReadAll -.
func (r *Repository) ReadAll(ctx context.Context, limit, offset uint64) ([]domain.Org, error) {
	return nil, nil
}

// ReadByID -.
func (r *Repository) ReadByID(ctx context.Context, id uint64) (domain.Org, error) {
	return domain.Org{}, nil
}

// ReadHistoryByID -.
func (r *Repository) ReadHistoryByID(ctx context.Context, id, limit, offset uint64) ([]domain.Org, error) {
	return nil, nil
}

// UpdateByID -.
func (r *Repository) UpdateByID(ctx context.Context, id uint64, o domain.Org) (domain.Org, error) {
	return domain.Org{}, nil
}

// SoftDeleteByID -.
func (r *Repository) SoftDeleteByID(ctx context.Context, id uint64) error {
	return nil
}
