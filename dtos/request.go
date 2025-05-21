package dtos

import (
	"github.com/PJNube/lib-models/rparams"
)

type RequestHandler struct {
	Pattern string
	Handler func(r rparams.Params) *APIResponse
}

type HandleRequestInput[P, B any] struct {
	Params P `json:"params,omitempty"`
	Body   B `json:"body,omitempty"`
}
