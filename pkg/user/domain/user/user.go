package user

import (
	"time"

	"github.com/gofrs/uuid"
)

const (
	RoleAdmin  string = "admin"
	RolePlayer string = "player"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	CreatedAt *time.Time
}
