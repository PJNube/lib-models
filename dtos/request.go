package dtos

import "github.com/PJNube/lib-models/qparams"

type RequestHandler struct {
	Pattern string
	Handler func(params qparams.Params, body any) *APIResponse
}

type HandleRequestInput[T any] struct {
	Params qparams.Params
	Body   T
}
