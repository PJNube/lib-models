package dtos

type DeviceInfo struct {
	GlobalUUID   string `json:"globalUuid"`
	Name         string `json:"name"`
	Architecture string `json:"architecture"`
}
