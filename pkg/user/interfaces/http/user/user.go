package user

import (
	"context"
	"net/http"

	"github.com/ghanto/jivala-core/pkg/user/domain/user"
	"github.com/gin-gonic/gin"
)

type userService interface {
	Register(ctx context.Context, dto user.User) (*user.User, error)
	GetByEmail(ctx context.Context, email string) (*user.User, error)
}

type UserHandler struct {
	us userService
}

func NewUsersHandler(us userService) *UserHandler {
	return &UserHandler{us: us}
}

func (h UserHandler) Create(c *gin.Context) {
	dto := CreateUserRequest{}
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.us.Register(c, user.User{
		Email: dto.Email,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (h UserHandler) Get(c *gin.Context) {
	dto := GetUserRequest{}
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usr, err := h.us.GetByEmail(c, dto.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserGetResponse{
		ID:        usr.ID,
		Email:     usr.Email,
		CreatedAt: usr.CreatedAt,
	})
}
