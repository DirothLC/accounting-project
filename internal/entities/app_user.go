package entities

import "Accounting/internal/utils"

type AppUser struct {
	login       string
	password    string
	accessToken string
	terminalID  int64
}

//go:noinline
func RegisterUser(login, password string) *AppUser {
	user := &AppUser{
		login:      login,
		password:   password,
		terminalID: utils.GenerateID(),
	}
	return user
}

func (a *AppUser) GetUserLogin() string {
	return a.login
}
