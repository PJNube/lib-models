package dtos

import "github.com/PJNube/lib-models/datatypes"

type ExtensionBatch struct {
	Type    datatypes.ExtensionBatchType `json:"type"`
	ID      string                       `json:"id"`
	Version string                       `json:"version"`
}
