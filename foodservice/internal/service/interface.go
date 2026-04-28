package service

import (
	m "mealmate/internal/model"
	"net/http"
	"net/url"
)

type Servicer interface {
	Create(*http.Request) ([]byte, int)
	Read(*http.Request) ([]byte, int)
	Update(*http.Request) ([]byte, int)
	Delete(*http.Request) ([]byte, int)
}

type ExFuncer interface {
	GetModelerID(map[string]any, string) (string, error)
	ReadRequestBody(*http.Request) ([]byte, error)
	TakeIDFromPath(*http.Request, string) (string, error)
	TakeIDFromBody(*http.Request, string) (string, error)
	NeedShow(m.Modeler, url.Values) bool
}
