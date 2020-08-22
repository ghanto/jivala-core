package user

import (
	"time"

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
