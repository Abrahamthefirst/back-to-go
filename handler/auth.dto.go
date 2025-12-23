package handler

type LoginResponseBody struct {
	AccessToken string `json:"access_token"`
}
type LoginResponse struct {
	Message string            `json:"message"`
	Body    LoginResponseBody `json:"body"`
	Error   string            `json:"error"`
}
type LoginRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=4,lte=20"`
}