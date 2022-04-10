package env

import "os"

var DevConfig = AppConfig{
	DBConfig: DBConfig{
		Driver:   "mysql",
		Host:     "localhost",
		Port:     "8080",
		Database: "trip_wishlist_db",
		User:     "mysql",
		Password: "mysql",
	},
}

var TestConfig = AppConfig{
	DBConfig: DBConfig{
		Driver:   "mysql",
		Host:     "http://ec2-3-83-138-94.compute-1.amazonaws.com",
		Port:     "8080",
		Database: "trip_wishlist_db",
		User:     "admin",
		Password: "minhasenha",
	},
}

func GetEnvConfig() AppConfig {
	environment := os.Getenv("GO_ENVIRONMENT")
	config := DevConfig
	if environment == "test" {
		config = TestConfig
	}
	return config
}
