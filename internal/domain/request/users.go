package request

type LoginUserRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
