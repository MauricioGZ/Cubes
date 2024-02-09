package dtos

type RegisterUser struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=8"`
	Name     string `form:"name"`
}

type LoginUser struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}
