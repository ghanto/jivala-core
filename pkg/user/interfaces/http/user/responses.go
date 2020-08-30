package user

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserGetResponse struct {
	ID        uuid.UUID
	Email     string
	CreatedAt *time.Time
}
