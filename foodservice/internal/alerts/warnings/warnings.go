package warnings

type Warning interface {
	Warning() string
}

type warningString struct {
	s string
}

func New(text string) Warning {
	return &warningString{text}
}

func (w *warningString) Warning() string {
	return w.s
}
