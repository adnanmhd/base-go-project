package service

type IService interface {
	GetToken() (*AuthResponse, error)
}
