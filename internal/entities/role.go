package entities

type Role struct {
	BaseGorm
	Name string `json:"name" validate:"required"`
}
