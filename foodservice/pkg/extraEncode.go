package pkg

import "encoding/json"

type ExEncoder interface {
	Deserialization([]byte, any) error
	Serialization(any) ([]byte, error)
}

type ExtraEncode struct {
}

func NewExtraEncode() *ExtraEncode {
	return &ExtraEncode{}
}

func (e *ExtraEncode) Deserialization(sliceByte []byte, pStruct any) error {
	return json.Unmarshal(sliceByte, pStruct)
}

func (e *ExtraEncode) Serialization(someStruct any) ([]byte, error) {
	return json.Marshal(someStruct)
}
