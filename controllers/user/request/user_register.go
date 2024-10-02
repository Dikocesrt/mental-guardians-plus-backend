package request

type UserRegister struct {
	Email     string `json:"email" form:"email" validate:"required,email"`
	Password  string `json:"password" form:"password" validate:"required"`
	FirstName string `json:"firstName" form:"firstName" validate:"required"`
	LastName  string `json:"lastName" form:"lastName" validate:"required"`
}