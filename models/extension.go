package models

type Extension struct {
	ID string `json:"id" gorm:"not null;primaryKey"`

	Resources []*Resource `json:"resources,omitempty" gorm:"foreignKey:ExtensionID;constraint:OnDelete:CASCADE"`
}
