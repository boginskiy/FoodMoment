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
		// Сбор необходимой инфы из Request
		start := time.Now()
		uri := r.RequestURI
		method := r.Method

		// Подготовка расширенного Response
		Exw := &ExResponseWriter{
			ResponseWriter: w,
			exData:         &ExData{},
		}

		// Передача запроса далее
		next.ServeHTTP(Exw, r)

		// Запись лога по Request
		m.Logg.RaiseInfo("Request:",
			"uri", uri,
			"method", method,
			"duration", time.Since(start))

		// Запись лога по Response
		m.Logg.RaiseInfo("Response:",
			"status", Exw.exData.status,
			"size", Exw.exData.size)
	})
}

func (m *Mdlware) WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if m.Auth.CheckAuthURL(r) {
			// Пропускаем urls, которые аворизуют/регистрируют
			next.ServeHTTP(w, r)

		} else {
			// Проводим аутентификацию
		}

	})
}
