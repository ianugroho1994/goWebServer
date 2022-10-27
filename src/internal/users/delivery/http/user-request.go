package http

type UserUpdateRequest struct {
	Name         string `json:"name" validate:"required"`
	Address      string `json:"address" validate:"required"`
	StartingDate string `json:"starting_date" validate:"required"`
	Email        string `json:"email" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
}
