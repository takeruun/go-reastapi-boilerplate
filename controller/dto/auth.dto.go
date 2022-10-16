package dto

type AuthSignInRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
