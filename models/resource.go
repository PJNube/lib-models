package models

import "gorm.io/datatypes"

type Resource struct {
	ID          string                      `json:"id" gorm:"not null;primaryKey"`
	ExtensionID string                      `json:"extensionId" gorm:"not null;primaryKey"`
	Description string                      `json:"description"`
	Items       datatypes.JSONSlice[string] `json:"items,omitempty"`

	Extension *Extension `json:"extension,omitempty" gorm:"foreignKey:ExtensionID"`
}
