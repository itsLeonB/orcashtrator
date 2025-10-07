package auth

type RegisterRequest struct {
	Email                string `validate:"required,email,min=3"`
	Password             string `validate:"required,eqfield=PasswordConfirmation,min=8"`
	PasswordConfirmation string `validate:"required"`
	VerificationURL      string
}

type InternalLoginRequest struct {
	Email    string `validate:"required,email,min=3"`
	Password string `validate:"required"`
}

type OAuthLoginRequest struct {
	Provider string `validate:"required,min=1"`
	Code     string `validate:"required,min=1"`
	State    string `validate:"required,min=1"`
}

type LoginResponse struct {
	Type  string
	Token string
}
