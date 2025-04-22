package qparams

import (
	"errors"
	"strconv"
)

type Params map[string][]string

func (p Params) Get(key string) string {
	vs := p[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

func (p Params) GetBool(key string) bool {
	vs := p[key]
	if len(vs) == 0 {
		return false
	}
	parseBool, _ := strconv.ParseBool(vs[0])
	return parseBool
}

func (p Params) GetInt(key string) int {
	vs := p[key]
	if len(vs) == 0 {
		return 0
	}
	intValue, _ := strconv.Atoi(vs[0])
	return intValue
}

func (p Params) Set(key, value string) {
	p[key] = []string{value}
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
