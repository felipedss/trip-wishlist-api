package dto

type FlightOfferResponse struct {
	Data []Data `json:"data"`
}

type Data struct {
	Type                  string         `json:"type"`
	NumberOfBookableSeats int            `json:"numberOfBookableSeats"`
	Itinerary             []Itinerary    `json:"itineraries"`
	Price                 Price          `json:"price"`
	PricingOptions        PricingOptions `json:"pricingOptions"`
}

type Itinerary struct {
	Duration string    `json:"duration"`
	Segment  []Segment `json:"segments"`
}

type Segment struct {
	Duration      string    `json:"duration"`
	Departure     Departure `json:"departure"`
	Arrival       Arrival   `json:"arrival"`
	NumberOfStops int       `json:"numberOfStops"`
}

type Departure struct {
	IataCode string `json:"iataCode"`
	Terminal string `json:"terminal"`
	At       string `json:"at"`
}

type Arrival struct {
	IataCode string `json:"iataCode"`
	Terminal string `json:"terminal"`
	At       string `json:"at"`
}

type Price struct {
	Currency string `json:"currency"`
	Total    string `json:"total"`
	Base     string `json:"base"`
	Fees     []struct {
		Amount string `json:"amount"`
		Type   string `json:"type"`
	} `json:"fees"`
	GrandTotal string `json:"grandTotal"`
}

type PricingOptions struct {
	FareType                []string `json:"fareType"`
	IncludedCheckedBagsOnly bool     `json:"includedCheckedBagsOnly"`
}
