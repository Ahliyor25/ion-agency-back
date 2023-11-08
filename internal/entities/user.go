package entities

// User — the basic structure of the user
// swagger:model
type User struct {
	BaseGorm
	FullName string `json:"full_name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	RoleId   uint   `json:"role_id" validate:"required"`

	RoleID   uint   `json:"-"`
	Password string `json:"password"`

	File []File `gorm:"many2many:user_files;"`
	Role Role   `json:"-"`
}
