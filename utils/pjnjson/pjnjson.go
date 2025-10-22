package pjnjson

import (
	"encoding/json"

	"gorm.io/datatypes"
)

func MarshalJson(jsonData datatypes.JSON) []byte {
	if jsonData == nil {
		jsonData = datatypes.JSON{}
	}
	mJsonData, _ := json.Marshal(jsonData)
	return mJsonData
}
