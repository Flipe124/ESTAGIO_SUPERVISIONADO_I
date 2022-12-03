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

type Setting struct {
	Api ApiSetting
	Database DatabaseSetting
}
