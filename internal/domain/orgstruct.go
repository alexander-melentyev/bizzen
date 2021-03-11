package domain

import "context"

// OrgStruct -.
type OrgStruct struct {
	Model
	OrgID    uint64 `db:"org_id" json:"orgId"`
	ParentID uint64 `db:"parent_id" json:"parentId"`
}

//go:generate mockgen -destination=../orgstruct/usecase/mock_repo_test.go -package=usecase_test . OrgStructRepository

// OrgStructRepository - org repository.
type OrgStructRepository interface {
	Create(ctx context.Context, orgStruct OrgStruct) error
}

//go:generate mockgen -destination=../orgstruct/delivery/http/mock_s_test.go -package=http_test . OrgStructUseCase

// OrgStructUseCase - biz usecases.
type OrgStructUseCase interface {
	Create(ctx context.Context, orgStruct OrgStruct) error
}
