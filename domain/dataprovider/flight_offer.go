package dataprovider

type FlightOfferClient interface {
	Get(filter map[string][]string, accept string, authorization string) ([]byte, error)
}
