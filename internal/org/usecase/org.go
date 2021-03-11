package usecase

import (
	"context"

	"github.com/alexander-melentyev/bizzen/internal/domain"
)

// UseCase -.
type UseCase struct {
	r domain.OrgRepository
}

var _ domain.OrgUseCase = (*UseCase)(nil)

// NewUseCase -.
func NewUseCase(r domain.OrgRepository) *UseCase {
	return &UseCase{
		r: r,
	}
}

// Create -.
func (uc *UseCase) Create(ctx context.Context, org domain.Org) error {
	return uc.r.Create(ctx, org)
}

// ReadAll -.
func (uc *UseCase) ReadAll(ctx context.Context, limit, offset uint64) ([]domain.Org, error) {
	return uc.r.ReadAll(ctx, limit, offset)
}

// ReadByID -.
func (uc *UseCase) ReadByID(ctx context.Context, id uint64) (domain.Org, error) {
	return uc.r.ReadByID(ctx, id)
}

// ReadHistoryByID -.
func (uc *UseCase) ReadHistoryByID(ctx context.Context, id, limit, offset uint64) ([]domain.Org, error) {
	return uc.r.ReadHistoryByID(ctx, id, limit, offset)
}

// UpdateByID -.
func (uc *UseCase) UpdateByID(ctx context.Context, id uint64, org domain.Org) (domain.Org, error) {
	return uc.r.UpdateByID(ctx, id, org)
}

// SoftDeleteByID -.
func (uc *UseCase) SoftDeleteByID(ctx context.Context, id uint64) error {
	return uc.r.SoftDeleteByID(ctx, id)
}
