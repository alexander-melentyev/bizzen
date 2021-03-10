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

const createQuery = `INSERT INTO org (name, creator, updater) VALUES (:name, :creator, :updater)`

// Create -.
func (r *Repository) Create(ctx context.Context, org domain.Org) error {
	org.Creator = ""
	org.Updater = ""

	_, err := r.conn.NamedExecContext(ctx, createQuery, org)

	return err
}

const readAllQuery = `SELECT * FROM org WHERE deleted_at IS NULL LIMIT $1 OFFSET $2`

// ReadAll -.
func (r *Repository) ReadAll(ctx context.Context, limit, offset uint64) ([]domain.Org, error) {
	var res []domain.Org

	return res, r.conn.SelectContext(ctx, &res, readAllQuery, limit, offset)
}

const readByIDQuery = `SELECT * FROM org WHERE deleted_at IS NULL AND id = $1`

// ReadByID -.
func (r *Repository) ReadByID(ctx context.Context, id uint64) (domain.Org, error) {
	var res domain.Org

	return res, r.conn.GetContext(ctx, &res, readByIDQuery, id)
}

// ReadHistoryByID -.
func (r *Repository) ReadHistoryByID(ctx context.Context, id, limit, offset uint64) ([]domain.Org, error) {
	return nil, nil
}

// UpdateByID -.
func (r *Repository) UpdateByID(ctx context.Context, id uint64, org domain.Org) (domain.Org, error) {
	return domain.Org{}, nil
}

// SoftDeleteByID -.
func (r *Repository) SoftDeleteByID(ctx context.Context, id uint64) error {
	return nil
}
