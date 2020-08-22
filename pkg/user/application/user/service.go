package user

import "github.com/ghanto/jivala-core/pkg/user/domain/user"

func NewService() {

}

type Service interface {
	Register(db interface{}, dto *user.CreateUser) (*user.User, error)
}
