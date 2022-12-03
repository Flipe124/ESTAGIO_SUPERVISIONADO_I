package settings

type ApiSetting struct {
	Port int
}

type DatabaseSetting struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

type Settings struct {
	Api      ApiSetting
	Database DatabaseSetting
}

func GetApiSetting() ApiSetting {
	return ApiSetting{
		Port: 8008,
	}
}

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

func GetSetting() Settings {
	return Settings{
		Api:      GetApiSetting(),
		Database: GetDatabaseSetting(),
	}
}
