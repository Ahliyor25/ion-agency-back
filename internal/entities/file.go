package entities

type File struct {
	BaseGorm
	File string `json:"file" validate:"required"`
}
