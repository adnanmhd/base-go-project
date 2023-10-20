package service

type Payload struct {
	Command      string `json:"command"`
	ResponseCode string `json:"response_code,omitempty"`
	Data         Data   `json:"data"`
}

type Data struct{}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
