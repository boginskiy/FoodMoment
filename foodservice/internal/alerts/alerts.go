package alerts

import (
	"encoding/json"
	"fmt"
	w "mealmate/internal/alerts/warnings"
	"net/http"
	"os"
	"time"
)

type Alerter interface {
	HandleAlert(*http.Request, any, int) ([]byte, int)
}

type Alert struct {
	Alert     string
	Code      int
	Message   string
	Path      string
	Timestamp time.Time
}

func NewAlert() *Alert {
	return &Alert{}
}

func (a *Alert) toMarshal() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Alert) bodyToMarshal() []byte {
	tmpBody, err := a.toMarshal()

	if err != nil {
		// Если не удалось сериализовать Error, пишем os.Stdout и короткое сообщение
		tmpStr := fmt.Sprintf(`{"%v message":"%v"}`, a.Alert, err)
		fmt.Fprintf(os.Stdout, "%v\n", tmpStr)
		tmpBody = []byte(tmpStr)
	}
	return tmpBody
}

func (a *Alert) preparBody(kind, mess, path string, statusCode int) {
	a.Timestamp = time.Now()
	a.Code = statusCode
	a.Message = mess
	a.Alert = kind
	a.Path = path
}

func (a *Alert) HandleAlert(req *http.Request, msg any, statusCode int) ([]byte, int) {
	var message string
	var kind string

	switch v := msg.(type) {
	case error:
		message = v.Error()
		kind = "error"
	case w.Warning:
		message = v.Warning()
		kind = "warning"
	default:
		message = fmt.Sprintf("%v\n", msg)
		kind = "unknown"
	}
	a.preparBody(kind, message, req.URL.Path, statusCode)
	return a.bodyToMarshal(), statusCode
}
