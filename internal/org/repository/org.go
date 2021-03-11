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

const createQuery = `INSERT INTO org (name, creator, updater)
VALUES (:name, :creator, :updater)`

// Create -.
func (r *Repository) Create(ctx context.Context, org domain.Org) (err error) {
	org.Creator = ""
	org.Updater = ""

	_, err = r.conn.NamedExecContext(ctx, createQuery, org)

	return
}

const readAllQuery = `SELECT *
FROM org
WHERE deleted_at IS NULL
LIMIT $1 OFFSET $2`

// ReadAll -.
func (r *Repository) ReadAll(ctx context.Context, limit, offset uint64) (res []domain.Org, err error) {
	err = r.conn.SelectContext(ctx, &res, readAllQuery, limit, offset)

	return
}

const readByIDQuery = `SELECT *
FROM org
WHERE deleted_at IS NULL AND id = $1`

// ReadByID -.
func (r *Repository) ReadByID(ctx context.Context, id uint64) (res domain.Org, err error) {
	err = r.conn.GetContext(ctx, &res, readByIDQuery, id)

	return
}

const readHistoryByIDQuery = `SELECT *
FROM org_history
WHERE deleted_at IS NULL AND id = $1
LIMIT $2 OFFSET $3`

// ReadHistoryByID -.
func (r *Repository) ReadHistoryByID(ctx context.Context, id, limit, offset uint64) (res []domain.Org, err error) {
	err = r.conn.SelectContext(ctx, &res, readHistoryByIDQuery, id, limit, offset)

	return
}

const updateByIDQuery = `UPDATE org
SET name = :name, updater = :updater
WHERE deleted_at IS NULL AND id = :id
RETURNING *`

// UpdateByID -.
func (r *Repository) UpdateByID(ctx context.Context, id uint64, org domain.Org) (domain.Org, error) {
	stmt, err := r.conn.PrepareNamedContext(ctx, updateByIDQuery)
	if err != nil {
		return domain.Org{}, err
	}

	org.ID = id
	org.Updater = ""

	var res domain.Org

	return res, stmt.GetContext(ctx, &res, org)
}

const deleteByIDQuery = `UPDATE org
SET deleter = :deleter, deleted_at = NOW()
WHERE deleted_at IS NULL AND id = :id
RETURNING *`

// SoftDeleteByID -.
func (r *Repository) SoftDeleteByID(ctx context.Context, id uint64) error {
	stmt, err := r.conn.PrepareNamedContext(ctx, deleteByIDQuery)
	if err != nil {
		return err
	}

	var org domain.Org

	org.ID = id
	org.Deleter.Valid = true
	org.Deleter.String = ""

	return stmt.GetContext(ctx, &org, org)
}
