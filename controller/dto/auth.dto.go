package dto

type AuthSignInRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthSignUpRequestDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
