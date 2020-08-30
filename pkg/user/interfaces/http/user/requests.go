package user

type CreateUserRequest struct {
	Email string `json:"email" binding:"required,email" example:"mariusz@dariusz.com"`
}

type GetUserRequest struct {
	Email string `json:"email" binding:"required,email" example:"mariusz@dariusz.com"`
}
