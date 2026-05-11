package logg

import (
	"encoding/json"
	"log/slog"
	"time"
)

type Prepar struct {
}

func (p *Prepar) SerializeToJSON(r slog.Record) ([]byte, error) {
	viewLog := map[string]interface{}{
		"time":    r.Time.Format(time.RFC3339),
		"level":   r.Level.String(),
		"message": r.Message,
		"attrs":   make(map[string]interface{}),
	}

	attrs := make(map[string]interface{})
	r.Attrs(func(a slog.Attr) bool {
		attrs[a.Key] = a.Value.Any()
		return true
	})

	if len(attrs) > 0 {
		viewLog["attrs"] = attrs
	}

	return json.Marshal(viewLog)
}
