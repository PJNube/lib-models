package dtos

import (
	"github.com/PJNube/lib-models/qparams"
	"github.com/PJNube/lib-models/rparams"
)

type RequestHandler struct {
	Pattern string
	Handler func(q qparams.Params, r rparams.Params, body any) *APIResponse
}

type HandleRequestInput[T any] struct {
	Params qparams.Params
	Body   T
}
