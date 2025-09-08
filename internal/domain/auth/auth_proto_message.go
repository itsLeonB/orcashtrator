package auth

type RegisterRequest struct {
	Email                string `validate:"required,email,min=3"`
	Password             string `validate:"required,eqfield=PasswordConfirmation,min=8"`
	PasswordConfirmation string `validate:"required"`
}

type LoginRequest struct {
	Email    string `validate:"required,email,min=3"`
	Password string `validate:"required"`
}

type LoginResponse struct {
	Type  string
	Token string
}
