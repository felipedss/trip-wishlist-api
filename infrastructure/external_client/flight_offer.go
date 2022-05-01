package external_client

import (
	"io/ioutil"
	"net/http"

	"github.com/felipedss/trip-wishlist-api/app/controller/rest"
	"github.com/felipedss/trip-wishlist-api/infrastructure/env"
)

func NewWFlightOfferClient() *NewWFlightOfferClientSt {
	return &NewWFlightOfferClientSt{}
}

type NewWFlightOfferClientSt struct {
}

func (p *NewWFlightOfferClientSt) Get(queryParams map[string][]string, accept string, authorization string) ([]byte, error) {
	urlPath, err := rest.BuildQueryParam(env.GetEnvConfig().AmadeusUrl, queryParams)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("GET", urlPath, nil)
	req.Header.Add("accept", accept)
	req.Header.Add("Authorization", authorization)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body, nil

}
