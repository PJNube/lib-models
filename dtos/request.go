package dtos

import (
	"github.com/PJNube/lib-models/qparams"
	"github.com/PJNube/lib-models/rparams"
)

type RequestHandler struct {
	Pattern string
	Handler func(r rparams.Params) *APIResponse
}

type HandleRequestInput[T any] struct {
	Params qparams.Params `json:"params,omitempty"`
	Body   T              `json:"body,omitempty"`
}
