package qparams

import (
	"errors"
	"fmt"
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
	} else if isString(vs) {
		return vs.(string)
	}

	return ""
}

func (p Params) GetBool(key string) bool {
	vs, ok := p[key]
	if !ok {
		return false
	}
	isString := isString(vs)
	if !isString {
		return false
	} else if isString {
		boolVal, err := strconv.ParseBool(vs.(string))
		if err != nil {
			return false
		}
		return boolVal
	}
	return false
}

func (p Params) GetInt(key string) int {
	vs, ok := p[key]
	if !ok {
		return 0
	}
	isString := isString(vs)
	if !isString {
		return 0
	} else if isString {
		intVal, err := strconv.Atoi(vs.(string))
		if err != nil {
			return 0
		}
		return intVal
	}
	return 0
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

func (p Params) RefreshParams() Params {
	var resultMap Params = make(Params)
	for k, v := range p {
		vT := reflect.TypeOf(v).Kind()
		if vT == reflect.Float32 || vT == reflect.Float64 || vT == reflect.Int || vT == reflect.Int32 || vT == reflect.Int64 || vT == reflect.Bool {
			resultMap[k] = fmt.Sprintf("%v", v)
		} else if vT == reflect.Slice || vT == reflect.Array {
			var listVal []string
			for _, val := range v.([]any) {
				listVal = append(listVal, fmt.Sprintf("%v", val))
			}
			resultMap[k] = listVal
		} else {
			resultMap[k] = v
		}
	}
	return resultMap
}
