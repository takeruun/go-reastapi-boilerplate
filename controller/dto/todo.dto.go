package dto

type TodoCreateRequestDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
