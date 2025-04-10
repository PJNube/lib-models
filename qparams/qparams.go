package qparams

import (
	"errors"
)

type Params struct {
	UUID              string `json:"uuid" form:"uuid"`
	ID                string `json:"id" form:"id"`
	Name              string `json:"name" form:"name"`
	ParentUUID        string `json:"parentUuid" form:"parentUuid"`
	Depth             uint8  `json:"depth" form:"depth"`
	Architecture      string `json:"architecture" form:"architecture"`
	WithHosts         bool   `json:"withHosts" form:"withHosts"`
	WithVersions      bool   `json:"withVersions" form:"withVersions"`
	WithNesting       bool   `json:"withNesting" form:"withNesting"`
	DownloadExtension bool   `json:"downloadExtension" form:"downloadExtension"`
	ForceUpload       bool   `json:"forceUpload" form:"forceUpload"`
	Force             bool   `json:"force" form:"force"`
	WithDashboard     bool   `json:"withDashboard" form:"withDashboard"`
	WithData          bool   `json:"withData" form:"withData"`
}

func ValidateParamsUUIDRequired(params *Params) error {
	return validateParams(params, validateUuidRequired)
}

func ValidateParamsNameRequired(params *Params) error {
	return validateParams(params, validateNameRequired)
}

func ValidateParamsIdRequired(params *Params) error {
	return validateParams(params, validateIdRequired)
}

func validateIdRequired(params *Params) error {
	if params.ID == "" {
		return errors.New("id cannot be empty")
	}
	return nil
}

func validateUuidRequired(request *Params) error {
	if request.UUID == "" {
		return errors.New("uuid cannot be empty")
	}
	return nil
}

func validateNameRequired(request *Params) error {
	if request.Name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

func validateParams(params *Params, validators ...func(params *Params) error) error {
	for _, validate := range validators {
		if err := validate(params); err != nil {
			return err
		}
	}
	return nil
}
