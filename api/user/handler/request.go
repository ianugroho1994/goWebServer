package handler

type ChangePasswordRequest struct {
	Password string `json:"password" validate:"required"`
}

type CreateUserRequest struct {
	Name     string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"Password" validate:"required"`
}
