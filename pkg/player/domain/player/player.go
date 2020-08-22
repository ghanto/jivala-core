package player

import "github.com/gofrs/uuid"

type Player struct {
	ID     uuid.UUID
	UserId *uuid.UUID
	name   string
	level  int
}
