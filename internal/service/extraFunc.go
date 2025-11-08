package service

import (
	"fmt"
	"io"
	m "mealmate/internal/model"
	"mealmate/pkg"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type ExtraFunc struct {
	exEncoder pkg.ExEncoder
}

func NewExtraFunc(exencd pkg.ExEncoder) *ExtraFunc {
	return &ExtraFunc{exEncoder: exencd}
}

func (ex *ExtraFunc) queryParamsGoToLower(params url.Values) url.Values {
	max := len(params)
	for key, value := range params {
		params[strings.ToLower(key)] = value

		if len(params) > max {
			delete(params, key)
		}
	}
	return params
}

func (ex *ExtraFunc) NeedShow(modeler m.Modeler, queryParams url.Values) bool {
	// Go to lower key for good compare continue ...
	queryParams = ex.queryParamsGoToLower(queryParams)

	// Get fields of struct of interface of Modeler
	_value := reflect.ValueOf(modeler)
	_type := reflect.TypeOf(modeler)
	numFields := _type.NumField()

	var flagCheck uint8

	// Go to every field of struct
	for i := 0; i < numFields; i++ {
		field := _type.Field(i)
		name := field.Name
		kind := field.Type.Kind()
		value := _value.FieldByName(name).Interface()

		// Check in field in params
		paramStr := queryParams.Get(strings.ToLower(name))
		flagCheck = 0

		// If the name is absent, go to continue
		if paramStr == "" {
			continue
		}

		// Filtering only easy kind of paramStr. Check only ==
		switch kind {

		case reflect.Int:
			paramInt, err := strconv.Atoi(paramStr)
			if err != nil {
				fmt.Fprintf(os.Stdout, "%v\n", err)
			}
			// Check Int
			if value.(int) == paramInt {
				flagCheck = 1 << 0
			}

		case reflect.String:
			// Check String
			if strings.EqualFold(value.(string), paramStr) {
				flagCheck = 1 << 0
			}

		case reflect.Float64:
			paramFlt, err := strconv.ParseFloat(paramStr, 64)
			if err != nil {
				fmt.Fprintf(os.Stdout, "%v\n", err)
			}
			// Check Float
			if value.(float64) == paramFlt {
				flagCheck = 1 << 0
			}
		}

		// Common Check
		if flagCheck == uint8(0) {
			return false
		}
	}
	return true
}

func (ex *ExtraFunc) ReadRequestBody(req *http.Request) ([]byte, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (f *ExtraFunc) GetModelerID(deserializedModeler map[string]any, id string) (string, error) {
	value, ok := deserializedModeler[id]
	if !ok {
		return "", notIdFieldErr
	}
	valueStr, ok := value.(string)
	if !ok {
		return "", notValidIdErr
	}
	return valueStr, nil
}

func (ex *ExtraFunc) TakeIDFromPath(req *http.Request, id string) (string, error) {
	queryParams := req.URL.Query()
	// Go to lower key for good compare continue ...
	queryParams = ex.queryParamsGoToLower(queryParams)
	name := queryParams.Get(strings.ToLower(id))
	if name == "" {
		return "", notIdFieldErr
	}
	return name, nil
}

func (ex *ExtraFunc) TakeIDFromBody(req *http.Request, id string) (string, error) {
	body, err := ex.ReadRequestBody(req)
	if err != nil {
		return "", err
	}

	var modelerForDelete map[string]any
	err = ex.exEncoder.Deserialization(body, &modelerForDelete)
	if err != nil {
		return "", err
	}

	// Check that is ID, like 'Name'
	value, err := ex.GetModelerID(modelerForDelete, id)
	if err != nil {
		return "", err
	}
	return value, nil
}
