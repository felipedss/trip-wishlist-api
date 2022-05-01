package service

import (
	"github.com/felipedss/trip-wishlist-api/domain/dataprovider"
)

type FlightOfferService interface {
	Get(queryParams map[string][]string, accept string, authorization string) ([]byte, error)
}

type FlightOfferServiceSt struct {
	flightOfferClient dataprovider.FlightOfferClient
}

func NewFlightOfferService(flightOfferClient dataprovider.FlightOfferClient) *FlightOfferServiceSt {
	return &FlightOfferServiceSt{
		flightOfferClient: flightOfferClient,
	}
}

func (flightOfferServiceSt *FlightOfferServiceSt) Get(queryParams map[string][]string, accept string, authorization string) ([]byte, error) {
	return flightOfferServiceSt.flightOfferClient.Get(queryParams, accept, authorization)
}
