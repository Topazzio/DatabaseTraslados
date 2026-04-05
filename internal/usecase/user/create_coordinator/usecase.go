package create_coordinator

import (
	"context"

	domain "eventra/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
}

type UseCase struct {
	repository Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{repository: repo}
}

func (u *UseCase) Execute(ctx context.Context, user domain.User) (domain.User, error) {
	user.Role = domain.CoordinatorRole

	return u.repository.Create(ctx, user)
}
