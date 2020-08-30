package user

import (
	"fmt"
	"net/http"

	"github.com/ghanto/jivala-core/pkg/user/domain/user"
)

type userService interface {
	Register(dto *CreateUser) (*user.User, error)
}

type Handler struct {
	us userService
}

func NewUsersHandler(us userService) Handler {
	return Handler{us: us}
}

func (p *Handler) UsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("/users")
	w.WriteHeader(http.StatusOK)
}
