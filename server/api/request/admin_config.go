package request

type AdminConfigList struct {
	Group string `form:"group" binding:"required"`
}

type AdminConfigUpdate struct {
	Group string         `json:"group" binding:"required"`
	Data  []*AdminConfig `json:"data"`
}

type AdminConfig struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
