package dtos

type Role struct {
	Name  string `json:"name"`
	Topic string `json:"topic"`
}

func (r *Role) ToStringArray() []string {
	return []string{r.Name, r.Topic}
}
