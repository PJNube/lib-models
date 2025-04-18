package qparams

import (
	"errors"
)

type Params map[string][]string

func (p Params) Get(key string) string {
	vs := p[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
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
