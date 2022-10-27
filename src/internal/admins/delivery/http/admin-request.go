package http

type AdminUpdateRequest struct {
	Name         string `json:"name" validate:"required"`
	Address      string `json:"address" validate:"required"`
	StartingDate string `json:"starting_date" validate:"required"`
	Email        string `json:"email" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required,numeric,startswith=62"`
}

type AdminRegisterRequest struct {
	Name         string `json:"name" validate:"required"`
	Address      string `json:"address" validate:"required"`
	StartingDate string `json:"starting_date" validate:"required"`
	Email        string `json:"email" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required,numeric,startswith=62"`
	Password string `json:"password" validate:"required"`
}