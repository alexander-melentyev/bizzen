package domain

import (
	"context"
	"database/sql"
	"time"
)

// Org - database table structure sorted by types.
type Org struct {
	ID        uint64         `db:"id" json:"id"`
	Name      string         `db:"name" json:"name"`
	Creator   string         `db:"creator" json:"creator"`
	Updater   string         `db:"updater" json:"updater"`
	CreatedAt time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time      `db:"updated_at" json:"updatedAt"`
	Deleter   sql.NullString `db:"deleter" json:"deleter"`
	DeletedAt sql.NullTime   `db:"deleted_at" json:"deletedAt"`
}

//go:generate mockgen -destination=../org/usecase/mock_test.go -package=usecase_test . OrgRepository

// OrgRepository - org repository.
type OrgRepository interface {
	Create(ctx context.Context, org Org) error
	ReadAll(ctx context.Context, limit, offset uint64) ([]Org, error)
	ReadByID(ctx context.Context, id uint64) (Org, error)
	ReadHistoryByID(ctx context.Context, id, limit, offset uint64) ([]Org, error)
	UpdateByID(ctx context.Context, id uint64, org Org) (Org, error)
	SoftDeleteByID(ctx context.Context, id uint64) error
}

//go:generate mockgen -destination=../org/delivery/http/mock_test.go -package=http_test . OrgUseCase

// OrgUseCase - biz usecases.
type OrgUseCase interface {
	Create(ctx context.Context, org Org) error
	ReadAll(ctx context.Context, limit, offset uint64) ([]Org, error)
	ReadByID(ctx context.Context, id uint64) (Org, error)
	ReadHistoryByID(ctx context.Context, id, limit, offset uint64) ([]Org, error)
	UpdateByID(ctx context.Context, id uint64, org Org) (Org, error)
	SoftDeleteByID(ctx context.Context, id uint64) error
}
