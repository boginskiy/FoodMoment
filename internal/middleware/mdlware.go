package middleware

import (
	"mealmate/cmd/config"
	"mealmate/internal/auth"
	"mealmate/internal/logg"
	"net/http"
	"time"
)

type Mdlware struct {
	Cfg  config.Config
	Logg logg.Logger
	Auth auth.Author
}

func NewMdlware(config config.Config, logger logg.Logger, auth auth.Author) *Mdlware {
	return &Mdlware{Cfg: config, Logg: logger, Auth: auth}
}

func (m *Mdlware) WithLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Сбор необходимой инфы
		start := time.Now()
		uri := r.RequestURI
		method := r.Method

		// Передача запроса далее
		next.ServeHTTP(w, r)

		// Запись лога
		m.Logg.RaiseInfo("Request:",
			"uri", uri,
			"method", method,
			"duration", time.Since(start))
	})
}

func (m *Mdlware) WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Пропускаем urls, которые аворизуют/регистрируют
		if m.Auth.CheckAuthURL(r) {
			next.ServeHTTP(w, r)
		}

		//

		// next.ServeHTTP(w, r)
	})
}
