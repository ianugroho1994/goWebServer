package domain

type Organization struct {
	Name string `json:"name" validate:"required" gorm:"unique"`
}
