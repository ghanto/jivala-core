package user

import (
	"context"

	"github.com/ghanto/jivala-core/pkg/user/domain/user"
)

type userRepo interface {
	Create(ctx context.Context, dto *user.User) (*user.User, error)
	GetByEmail(ctx context.Context, email string) (*user.User, error)
}
