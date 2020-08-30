package user

import (
	"context"

	"github.com/ghanto/jivala-core/pkg/user/domain/user"
)

type UserService interface {
	Register(ctx context.Context, dto user.User) (*user.User, error)
	GetByEmail(ctx context.Context, email string) (*user.User, error)
}

type Service struct {
	repo userRepo
}

func NewService(ur userRepo) Service {
	return Service{repo: ur}
}

func (s *Service) Register(ctx context.Context, dto *user.User) (*user.User, error) {
	password := "tbd"
	dto.Password = password

	u, err := s.repo.Create(ctx, dto)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	u, err := s.repo.GetByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return u, nil

}
