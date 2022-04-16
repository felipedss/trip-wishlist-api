package env

import "os"

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

func GetEnvConfig() AppConfig {
	environment := os.Getenv("GO_ENVIRONMENT")
	config := DevConfig
	if environment == "test" {
		config = TestConfig
	}
	return config
}
