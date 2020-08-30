package user

import (
	"context"

	"github.com/ghanto/jivala-core/pkg/user/domain/user"
)

type UserService interface {
	Register(ctx context.Context, dto user.User) (*user.User, error)
	GetByEmail(ctx context.Context, email string) (*user.User, error)
}

type service struct {
	repo userRepo
}

func NewService(ur userRepo) UserService {
	return &service{repo: ur}
}

func (s *service) Register(ctx context.Context, dto user.User) (*user.User, error) {
	password := "tbd"
	dto.Password = password

	u, err := s.repo.Create(ctx, &dto)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *service) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	u, err := s.repo.GetByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return u, nil

}
