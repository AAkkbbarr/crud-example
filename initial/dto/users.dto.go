package dto

type CreateUsersDTO struct {
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    float64 `json:"email"`
}

type UpdateUsersDTO struct {
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    float64 `json:"email"`
}
