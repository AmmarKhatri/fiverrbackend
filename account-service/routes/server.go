package routes

import (
	"account-service/account"
)

type AccountServer struct {
	account.AccountServer
}

func NewAccountServer() *AccountServer {
	//define connections to your databases here
	return &AccountServer{}
}
