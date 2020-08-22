package user

import (
	"github.com/gofrs/uuid"
)

const (
	RoleAdmin  string = "admin"
	RolePlayer string = "player"
)

type User struct {
	ID       uuid.UUID
	Login    string
	Email    string
	Password string
}

type CreateUser struct {
	Login string
	Email string
}
