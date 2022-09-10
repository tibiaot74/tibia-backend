package requests

type GenerateTokenRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type GenerateTokenResponse struct {
	Token string `json:"token"`
}
