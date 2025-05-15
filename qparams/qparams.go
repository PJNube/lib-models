package qparams

import (
	"encoding/json"
	"net/url"
)

type Params struct {
	Name              string `json:"name"`
	Depth             int    `json:"depth"`
	Address           string `json:"address"`
	ForceUpload       bool   `json:"force_upload"`
	DownloadExtension bool   `json:"download_extension"`
	Architecture      string `json:"architecture"`
	ParentUuid        string `json:"parent_uuid"`
	WithNesting       bool   `json:"with_nesting"`
	WithHosts         bool   `json:"with_hosts"`
	WithVersions      bool   `json:"with_versions"`
	WithData          bool   `json:"with_data"`
	WithDashboard     bool   `json:"with_dashboard"`
}

func GetParam(params url.Values) (*Params, error) {
	result := new(Params)
	err := json.Unmarshal([]byte(params["params"][0]), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
