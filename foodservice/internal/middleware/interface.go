package middleware

import "net/http"

type Middleware interface {
	WithLogger(next http.Handler) http.Handler
	WithAuth(next http.Handler) http.Handler
}
