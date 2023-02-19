package user

type AuthResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
