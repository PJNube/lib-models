package dtos

import "github.com/pjnube/lib-models/qparams"

type HandleRequestInput[T any] struct {
	Params *qparams.Params
	Body   T
}
