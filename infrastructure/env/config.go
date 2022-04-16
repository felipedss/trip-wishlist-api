package env

type AppConfig struct {
	DBConfig
	ExternalClientConfig
}

type DBConfig struct {
	Driver   string
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

type ExternalClientConfig struct {
	Url string
}
