package settings

// DatabaseSetting é a estrutra que guarda as informações de configurações do banco de dados.
type DatabaseSetting struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

// GetDatabaseSetting retorna uma estrutura das configurações do banco de dados.
func GetDatabaseSetting() DatabaseSetting {
	return DatabaseSetting{
		Driver:   "mysql",
		User:     "root",
		Password: "root",
		Host:     "172.17.0.2",
		Port:     3306,
		Name:     "openfinance",
	}
}
