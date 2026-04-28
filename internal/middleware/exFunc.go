package middleware

import "net/http"

type ExData struct {
	status int
	size   int
}

type ExResponseWriter struct {
	// Встраиваем оригинальный http.ResponseWriter
	http.ResponseWriter
	exData *ExData
}

// Переопределеяем методы http.ResponseWriter для получения доп инфы из Response

func (exR *ExResponseWriter) Write(b []byte) (int, error) {
	// записываем ответ, используя оригинальный http.ResponseWriter
	size, err := exR.ResponseWriter.Write(b)
	exR.exData.size += size
	return size, err
}

func (exR *ExResponseWriter) WriteHeader(statusCode int) {
	exR.ResponseWriter.WriteHeader(statusCode)
	exR.exData.status = statusCode
}
