package dtos

import "github.com/PJNube/lib-models/qparams"

type HandleRequestInput[T any] struct {
	Params *qparams.Params
	Body   T
}
