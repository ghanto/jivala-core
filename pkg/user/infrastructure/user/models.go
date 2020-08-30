package user

import (
	"time"

	"github.com/ghanto/jivala-core/pkg/user/domain/user"
	"github.com/gofrs/uuid"
)

type UserDB struct {
	Id        uuid.UUID `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type CreateUserDb struct {
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u *UserDB) ToModel() *user.User {
	return &user.User{
		ID:        u.Id,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: &u.CreatedAt,
	}
}
