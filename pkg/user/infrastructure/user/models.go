package user

import (
	"time"

	"github.com/ghanto/jivala-core/pkg/user/domain/user"
	"github.com/gofrs/uuid"
)

type UserDB struct {
	Id        uuid.UUID `db:"id"`
	Login     string    `db:"login"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type CreateUserDb struct {
	Login     string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u *UserDB) ToModel() *user.User {
	return &user.User{
		ID:       u.Id,
		Login:    u.Login,
		Email:    u.Email,
		Password: u.Password,
	}
}
