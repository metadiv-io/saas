package types

type Api struct {
	Method string  `json:"method"`
	Path   string  `json:"path"`
	UUID   string  `json:"uuid"`
	Credit float64 `json:"credit"`
}

func (api *Api) Tag() string {
	return api.Method + ":" + api.Path
}
