package user

import (
	"context"

	"github.com/ghanto/jivala-core/pkg/user/domain/user"
)

type MockRepo struct {
	userToReturn *user.User
	errToReturn  error
}

func NewMockRepo(err error, u *user.User) *MockRepo {
	return &MockRepo{errToReturn: err, userToReturn: u}
}

func (s *MockRepo) Create(ctx context.Context, dto *user.User) (*user.User, error) {
	return s.userToReturn, s.errToReturn
}

func (s *MockRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	return s.userToReturn, s.errToReturn
}
