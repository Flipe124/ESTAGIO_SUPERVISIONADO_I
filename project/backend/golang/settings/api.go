package settings

// ApiSetting é a estrutra que guarda as informações de configurações da API.
type ApiSetting struct {
	Port int
}

// GetApiSetting retorna uma estrutura das configurações da API.
func GetApiSetting() ApiSetting {
	return ApiSetting{
		Port: 8008,
	}
}
