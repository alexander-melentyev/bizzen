package domain

import "context"

// Org -.
type Org struct {
	Model
	Name string `db:"name" json:"name"`
}

//go:generate mockgen -destination=../org/usecase/mock_repo_test.go -package=usecase_test . OrgRepository

// OrgRepository - org repository.
type OrgRepository interface {
	Create(ctx context.Context, o Org) error
	ReadAll(ctx context.Context, limit, offset uint64) ([]Org, error)
	ReadByID(ctx context.Context, id uint64) (Org, error)
	ReadHistoryByID(ctx context.Context, id, limit, offset uint64) ([]Org, error)
	UpdateByID(ctx context.Context, id uint64, o Org) (Org, error)
	SoftDeleteByID(ctx context.Context, id uint64) error
}

//go:generate mockgen -destination=../org/delivery/http/mock_s_test.go -package=http_test . OrgUseCase

// OrgUseCase - biz usecases.
type OrgUseCase interface {
	Create(ctx context.Context, o Org) error
	ReadAll(ctx context.Context) ([]Org, error)
	ReadByID(ctx context.Context, id uint64) (Org, error)
	ReadHistoryByID(ctx context.Context, id, limit, offset uint64) ([]Org, error)
	UpdateByID(ctx context.Context, id uint64, o Org) (Org, error)
	SoftDeleteByID(ctx context.Context, id uint64) error
}
