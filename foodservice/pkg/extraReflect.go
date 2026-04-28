package pkg

import (
	"reflect"
	"strings"
)

type ExReflecter interface {
	CrossUpdateStructs(currStruct any, newMap map[string]any) any
	ShowdownFullStruct(any) map[string]map[string]any
	ShowdownFieldsStruct(any) (fieldsList []string)
}

type ExtraReflect struct {
}

func NewExtraReflect() *ExtraReflect {
	return &ExtraReflect{}
}

func (t *ExtraReflect) ShowdownFullStruct(someStruct any) map[string]map[string]any {
	infoStruct := map[string]map[string]any{}

	_type := reflect.TypeOf(someStruct)
	_value := reflect.ValueOf(someStruct)

	if _type.Kind() == reflect.Struct {
		numFields := _type.NumField()

		// Проходим по каждому полю структуры
		for i := 0; i < numFields; i++ {
			field := _type.Field(i)
			name := field.Name
			kind := field.Type.Kind()
			value := _value.FieldByName(name).Interface()

			infoStruct[name] = map[string]any{"kind": kind.String(), "value": value, "count": numFields}
		}
	}
	return infoStruct
}

func (t *ExtraReflect) ShowdownFieldsStruct(someStruct any) (fieldsList []string) {
	_type := reflect.TypeOf(someStruct)

	if _type.Kind() == reflect.Struct {
		numFields := _type.NumField()
		fieldsList = make([]string, numFields)

		// Собираем поля структуры
		for i := 0; i < numFields; i++ {
			field := _type.Field(i)
			fieldsList[i] = strings.ToLower(field.Name)
		}
	}
	return fieldsList
}

func (t *ExtraReflect) CrossUpdateStructs(currStruct any, newMap map[string]any) any {
	// Data of currStruct
	_type := reflect.TypeOf(currStruct).Elem()
	_value := reflect.ValueOf(currStruct).Elem()

	if _type.Kind() == reflect.Struct {

		for i := 0; i < _type.NumField(); i++ {
			field := _type.Field(i)
			name := field.Name
			kind := field.Type.Kind()

			// If field is in Map for update, update
			if tmpValue, ok := newMap[name]; ok {
				newValue := reflect.ValueOf(tmpValue)
				updateField := _value.FieldByName(name)

				// Check in before wraiting
				if updateField.CanSet() && kind == newValue.Kind() {
					updateField.Set(newValue)
				}
			}
		}
	}
	return currStruct
}
