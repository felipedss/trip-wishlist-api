package env

var DevConfig = AppConfig{
	DBConfig: DBConfig{
		Driver:   "mysql",
		Host:     "localhost",
		Port:     "8080",
		Database: "trip_wishlist_db",
		User:     "mysql",
		Password: "mysql",
	},
	ExternalClientConfig: ExternalClientConfig{
		AmadeusUrl: "https://test.api.amadeus.com/v2/shopping/flight-offers",
	},
}
