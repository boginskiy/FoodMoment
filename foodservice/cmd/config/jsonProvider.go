package config

import (
	"context"
	"log"

	"github.com/boginskiy/FoodMoment/foodservice/pkg"
)

type JSONProvider struct {
	Getter pkg.FlagEnvGetter
	Reader pkg.FileReader
	data   map[string]interface{}
}

func NewJSONProvider(ctx context.Context, getter pkg.FlagEnvGetter, reader pkg.FileReader) *JSONProvider {
	tmpPr := &JSONProvider{
		Getter: getter,
		Reader: reader,
		data:   map[string]interface{}{},
	}

	// Const variables.
	path := tmpPr.choosePath(nameOfVarPathCfgFile, defaultPathCfgFile)

	// Read config.json.
	err := tmpPr.Reader.Deserialization(path, &tmpPr.data)
	if err != nil {
		log.Printf("failed to deserialize config.json: %s", err.Error())
	}

	return tmpPr
}

func (j *JSONProvider) Load(key string) (any, bool) {
	if val, ok := j.data[key]; ok {
		return val, true
	}
	return nil, false
}

func (j *JSONProvider) choosePath(key, defaultVal string) string {
	if path := j.Getter.GetValueFromENV(key); path != "" {
		return path
	}
	if path := j.Getter.GetValueFromCLI(key); path != "" {
		return path
	}
	return defaultVal
}
