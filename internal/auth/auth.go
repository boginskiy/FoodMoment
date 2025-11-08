package auth

import (
	"fmt"
	"mealmate/cmd/config"
	"mealmate/internal/logg"
	"net/http"
)

type Auth struct {
	Cfg  config.Config
	Logg logg.Logger
}

func NewAuth(config config.Config, logger logg.Logger) *Auth {
	return &Auth{
		Cfg:  config,
		Logg: logger,
	}
}

func (a *Auth) Authentication() {

}

func (a *Auth) CheckAuthURL(req *http.Request) bool {
	fmt.Println(req.URL)
	return true
}
