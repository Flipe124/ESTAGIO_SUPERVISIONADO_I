package aarm

import "strconv"

// Api é a estrutura que representa a api.
type Api struct {
	Port int
}

// GetPortString retorna o valor da porta convertido em string.
func (api *Api) GetPortString() string {
	return strconv.Itoa(api.Port)
}

// NewApi retorna o potência de uma nova instância de uma estrutura Api.
func NewApi(port int) *Api {
	return &Api{
		Port: port,
	}
}
