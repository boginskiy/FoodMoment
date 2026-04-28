package auth

import (
	"mealmate/cmd/config"
	"mealmate/internal/logg"
	"net/http"
	"strings"
)

var ALLOWED_PATHS = map[string]struct{}{"register": struct{}{}, "login": struct{}{}}

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
	tmpURL := strings.Split(req.URL.String(), "/")
	tmpPartOfURL := tmpURL[len(tmpURL)-1]

	if _, ok := ALLOWED_PATHS[tmpPartOfURL]; ok {
		return true
	}
	return false
}
