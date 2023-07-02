package routes

import "auth-service/auth"

type AuthServer struct {
	auth.AuthServer
}

func NewAuthServer() *AuthServer {
	//define connections to your databases here
	return &AuthServer{}
}
