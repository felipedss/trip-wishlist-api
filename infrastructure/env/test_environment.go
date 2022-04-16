package env

import "os"

var TestConfig = AppConfig{
	DBConfig: DBConfig{
		Driver:   "mysql",
		Host:     "http://ec2-3-83-138-94.compute-1.amazonaws.com",
		Port:     "8080",
		Database: "trip_wishlist_db",
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	},
	ExternalClientConfig: ExternalClientConfig{
		Url: "https://test.api.amadeus.com/v2/shopping/flight-offers",
	},
}
