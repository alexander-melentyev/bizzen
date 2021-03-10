package usecase

import (
	"context"

	"github.com/alexander-melentyev/bizzen/internal/domain"
)

// UseCase -.
type UseCase struct {
	repo domain.OrgRepository
}

var _ domain.OrgUseCase = (*UseCase)(nil)

// NewUseCase -.
func NewUseCase(repo domain.OrgRepository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

// Create -.
func (uc *UseCase) Create(ctx context.Context, o domain.Org) error {
	return nil
}

// ReadAll -.
func (uc *UseCase) ReadAll(ctx context.Context, limit, offset uint64) ([]domain.Org, error) {
	return nil, nil
}

// ReadByID -.
func (uc *UseCase) ReadByID(ctx context.Context, id uint64) (domain.Org, error) {
	return domain.Org{}, nil
}

// ReadHistoryByID -.
func (uc *UseCase) ReadHistoryByID(ctx context.Context, id, limit, offset uint64) ([]domain.Org, error) {
	return nil, nil
}

// UpdateByID -.
func (uc *UseCase) UpdateByID(ctx context.Context, id uint64, o domain.Org) (domain.Org, error) {
	return domain.Org{}, nil
}

// SoftDeleteByID -.
func (uc *UseCase) SoftDeleteByID(ctx context.Context, id uint64) error {
	return nil
}
