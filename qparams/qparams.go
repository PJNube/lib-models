package qparams

import (
	"encoding/json"
	"errors"
	"net/url"
	"reflect"
	"strconv"
)

type Params map[string]any

func (p Params) Get(key string) string {
	vs, ok := p[key]
	if !ok {
		return ""
	}
	if !isString(vs) {
		return ""
	} else {
		return vs.(string)
	}
}

func (p Params) GetBool(key string) bool {
	vs, ok := p[key]
	if !ok {
		return false
	}
	if !isString(vs) {
		return false
	} else {
		boolVal, err := strconv.ParseBool(vs.(string))
		if err != nil {
			return false
		}
		return boolVal
	}
}

func (p Params) GetInt(key string) int {
	vs, ok := p[key]
	if !ok {
		return 0
	}
	if !isString(vs) {
		return 0
	} else {
		intVal, err := strconv.Atoi(vs.(string))
		if err != nil {
			return 0
		}
		return intVal
	}
}

func (p Params) Set(key string, value any) {
	p[key] = value
}

func (p Params) Del(key string) {
	delete(p, key)
}

func (p Params) Has(key string) bool {
	_, ok := p[key]
	return ok
}

func (p Params) Validate(key string) error {
	if !p.Has(key) {
		return errors.New(key + " is required")
	}
	return nil
}

func isString(value any) bool {
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		return false
	} else if reflect.TypeOf(value).Kind() == reflect.String {
		return true
	}
	return false
}

func GetParams(params url.Values) Params {
	result := make(Params)
	for key, value := range params {
		if len(value) == 1 {
			result[key] = value[0]
		} else {
			result[key] = value
		}
	}
	return result
}

func Parse[T any](params []byte) (*T, error) {
	result := new(T)
	err := json.Unmarshal(params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
