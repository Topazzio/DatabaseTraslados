package create

import (
	"context"
	"log"
	"strings"

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
	user.Name = strings.TrimSpace(user.Name)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Role = domain.CompanyRole

	if err := user.Validate(); err != nil {
		log.Printf("[Layer:UseCase][error_message:%s][request_body:%+v]", err.Error(), user)
		return domain.User{}, err
	}

	result, err := u.repository.Create(ctx, user)
	if err != nil {
		log.Printf("[Layer:UseCase][error_message:%s][request_body:%+v]", err.Error(), user)
		return domain.User{}, err
	}

	return result, nil
}
