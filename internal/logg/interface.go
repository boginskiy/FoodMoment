package logg

type Logger interface {
	RaiseInfo(msg string, keysAndValues any)
	RaiseWarning(msg string, keysAndValues any)
	RaiseError(msg string, err error)
	RaiseFatal(msg string, err error)
}
