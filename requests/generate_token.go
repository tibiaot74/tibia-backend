package requests

type GenerateTokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GenerateTokenResponse struct {
	Token string `json:"token"`
}
