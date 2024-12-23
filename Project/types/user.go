package types

type UserResponse struct {
	Id       int     `json:"id"`
	UserName string  `json:"username"`
	Phone    string  `json:"phone"`
	Money    float64 `json:"money"`
}
