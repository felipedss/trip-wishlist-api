package env

type AppConfig struct {
	DBConfig
}

type DBConfig struct {
	Driver   string
	Host     string
	Port     string
	Database string
	User     string
	Password string
}
