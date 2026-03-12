package request

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
