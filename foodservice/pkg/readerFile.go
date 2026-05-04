package pkg

import (
	"context"
	"encoding/json"
	"os"
)

type FileReader interface {
	Deserialization(path string, store any) error
}

type ReadFile struct {
}

func NewReadFile(ctx context.Context) *ReadFile {
	return &ReadFile{}
}

func (f *ReadFile) Deserialization(path string, store any) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, store)
	if err != nil {
		return err
	}
	return nil
}
