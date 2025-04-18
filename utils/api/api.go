package api

import (
	"errors"
	"github.com/PJNube/lib-models/dtos"
	"strings"
)

func BuildGetAPIResponse(data any, err error) *dtos.APIResponse {
	if err != nil {
		return GetServerError(err.Error())
	}
	return GetSuccess(200, data)
}

func BuildCreateAPIResponse(data any, err error) *dtos.APIResponse {
	if err != nil {
		return GetServerError(err.Error())
	}
	return GetSuccess(201, data)
}

func BuildDeleteAPIResponse(err error) *dtos.APIResponse {
	if err != nil {
		return GetServerError(err.Error())
	}
	return GetSuccess(204, nil)
}

func BuildNotFoundAPIResponse(message string) *dtos.APIResponse {
	return GetNotFoundError(message)
}

func BuildAPIParseError(err error) *dtos.APIResponse {
	msg := err.Error()
	return &dtos.APIResponse{
		Status:  "error",
		Code:    400,
		Message: &msg,
		Data:    nil,
	}
}

func GetSuccess(code int, data any) *dtos.APIResponse {
	return &dtos.APIResponse{
		Status: "success",
		Code:   code,
		Data:   data,
	}
}

func GetNotFoundError(message string) *dtos.APIResponse {
	return &dtos.APIResponse{
		Status:  "error",
		Code:    404,
		Message: &message,
		Data:    nil,
	}
}

func GetServerError(message string) *dtos.APIResponse {
	return &dtos.APIResponse{
		Status:  "error",
		Code:    500,
		Message: &message,
		Data:    nil,
	}
}

func ParseBody[T any](body any) (*T, error) {
	entity, ok := body.(*T)
	if !ok {
		return nil, errors.New("invalid body")
	}
	return entity, nil
}

func ResolveRouteParams(pattern, routes string) (map[string][]string, bool) {
	pTokens := strings.Split(pattern, ".")
	sTokens := strings.Split(routes, ".")

	if len(pTokens) != len(sTokens) {
		return nil, false
	}
	result := make(map[string][]string)

	for i := 0; i < len(pTokens); i++ {
		switch pTokens[i] {
		case "*":
			if i > 0 {
				key := pTokens[i-1]
				result[key] = []string{sTokens[i]}
			} else {
				result["*"] = []string{sTokens[i]}
			}
		default:
			if pTokens[i] != sTokens[i] {
				return nil, false
			}
		}
	}
	return result, true
}
