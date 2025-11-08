package logg

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
)

var NameFolder = "logs"

type Logg struct {
	Name string
	file *os.File
	zap  *zap.SugaredLogger
}

func NewLogg(fileName, level string) *Logg {
	pathToFile := fmt.Sprintf("./%s/%s", NameFolder, fileName)

	// Create folder and file
	f := MakeDirAndFile(NameFolder, pathToFile)

	// Configuration
	cfg := Config(pathToFile, level)

	// Create
	zapLogg, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	return &Logg{
		Name: fileName,
		zap:  zapLogg.Sugar(),
		file: f,
	}
}

func (l *Logg) Close() {
	defer l.file.Close()
	defer l.zap.Sync()
}

func (l *Logg) disassemblyAny(keysAndValues ...any) string {
	if len(keysAndValues) == 0 || len(keysAndValues)&1 == 1 {
		return ""
	}
	var buf bytes.Buffer
	for i := 0; i < len(keysAndValues); {
		buf.WriteString(fmt.Sprintf(" %v:%v", keysAndValues[i], keysAndValues[i+1]))
		i += 2
	}
	return buf.String()
}

func (l *Logg) RaiseInfo(msg string, keysAndValues ...any) {
	l.zap.Infoln(msg, keysAndValues)
}

func (l *Logg) RaiseWarning(msg string, keysAndValues ...any) {
	l.zap.Warnln(msg, keysAndValues)
}

func (l *Logg) RaiseError(msg string, err error) {
	if err == nil {
		l.zap.Errorf("%s", msg)
	} else {
		l.zap.Errorf("%s> %s", msg, err.Error())
	}
}

func (l *Logg) RaiseFatal(msg string, err error) {
	if err == nil {
		l.zap.Fatalf("%s", msg)
	} else {
		l.zap.Fatalf("%s> %s", msg, err.Error())
	}
}
