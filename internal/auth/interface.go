package auth

import "net/http"

type Author interface {
	CheckAuthURL(req *http.Request) bool
}
