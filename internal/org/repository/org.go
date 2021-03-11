package repository

import (
	"context"
	"fmt"

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

const createQuery = `INSERT INTO org (name, creator, updater)
VALUES (:name, :creator, :updater)`

// Create -.
func (r *Repository) Create(ctx context.Context, org domain.Org) error {
	org.Creator = ""
	org.Updater = ""

	_, err := r.conn.NamedExecContext(ctx, createQuery, org)
	if err != nil {
		return fmt.Errorf("failed to create org: %w", err)
	}

	return nil
}

const readAllQuery = `SELECT *
FROM org
WHERE deleted_at IS NULL
LIMIT $1 OFFSET $2`

// ReadAll -.
func (r *Repository) ReadAll(ctx context.Context, limit, offset uint64) ([]domain.Org, error) {
	var res []domain.Org

	if err := r.conn.SelectContext(ctx, &res, readAllQuery, limit, offset); err != nil {
		return nil, fmt.Errorf("failed to read all org: %w", err)
	}

	return res, nil
}

const readByIDQuery = `SELECT *
FROM org
WHERE deleted_at IS NULL AND id = $1`

// ReadByID -.
func (r *Repository) ReadByID(ctx context.Context, id uint64) (domain.Org, error) {
	var res domain.Org

	if err := r.conn.GetContext(ctx, &res, readByIDQuery, id); err != nil {
		return domain.Org{}, fmt.Errorf("failed to read org by ID: %w", err)
	}

	return res, nil
}

const readHistoryByIDQuery = `SELECT *
FROM org_history
WHERE deleted_at IS NULL AND id = $1
LIMIT $2 OFFSET $3`

// ReadHistoryByID -.
func (r *Repository) ReadHistoryByID(ctx context.Context, id, limit, offset uint64) ([]domain.Org, error) {
	var res []domain.Org

	if err := r.conn.SelectContext(ctx, &res, readHistoryByIDQuery, id, limit, offset); err != nil {
		return nil, fmt.Errorf("failed to read org history by ID: %w", err)
	}

	return res, nil
}

const updateByIDQuery = `UPDATE org
SET name = :name, updater = :updater
WHERE deleted_at IS NULL AND id = :id
RETURNING *`

// UpdateByID -.
func (r *Repository) UpdateByID(ctx context.Context, id uint64, org domain.Org) (domain.Org, error) {
	stmt, err := r.conn.PrepareNamedContext(ctx, updateByIDQuery)
	if err != nil {
		return domain.Org{}, fmt.Errorf("failed to create prepare statement for update org by ID: %w", err)
	}

	org.ID = id
	org.Updater = ""

	var res domain.Org

	if err := stmt.GetContext(ctx, &res, org); err != nil {
		return domain.Org{}, fmt.Errorf("failed to update org by ID: %w", err)
	}

	return res, nil
}

const deleteByIDQuery = `UPDATE org
SET deleter = :deleter, deleted_at = NOW()
WHERE deleted_at IS NULL AND id = :id
RETURNING *`

// SoftDeleteByID -.
func (r *Repository) SoftDeleteByID(ctx context.Context, id uint64) error {
	stmt, err := r.conn.PrepareNamedContext(ctx, deleteByIDQuery)
	if err != nil {
		return fmt.Errorf("failed to create prepare statement for soft delete org by ID: %w", err)
	}

	var org domain.Org

	org.ID = id
	org.Deleter.Valid = true
	org.Deleter.String = ""

	if err := stmt.GetContext(ctx, &org, org); err != nil {
		return fmt.Errorf("failed to soft delete org by ID: %w", err)
	}

	return nil
}
