package handlers

import (
	"mealmate/cmd/config"
	"mealmate/internal/logg"
	"net/http"
)

type AuthHdler struct {
	Cfg  config.Config
	Logg logg.Logger
}

func NewAuthHandler(config config.Config, logger logg.Logger) *AuthHdler {
	return &AuthHdler{Cfg: config, Logg: logger}
}

func (a *AuthHdler) ResetPass(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ResetPass"))
}

func (a *AuthHdler) Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register"))
}

func (a *AuthHdler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}

func (a *AuthHdler) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}
