package controller

import (
	"encoding/json"
	"net/http"

	"github.com/felipedss/trip-wishlist-api/app/controller/apierror"
	"github.com/felipedss/trip-wishlist-api/app/controller/dto"
	"github.com/felipedss/trip-wishlist-api/domain/service"
	"github.com/gin-gonic/gin"
)

type FlightOfferController interface {
	GetAll(c *gin.Context)
}

type FlightOfferControllerSt struct {
	flightOfferService service.FlightOfferService
}

func NewWFlightOfferController(flightOfferService service.FlightOfferService) *FlightOfferControllerSt {
	return &FlightOfferControllerSt{
		flightOfferService: flightOfferService,
	}
}

func (p *FlightOfferControllerSt) GetAll(c *gin.Context) {

	body, err := p.flightOfferService.Get(c.Request.URL.Query(), c.GetHeader("accept"), c.GetHeader("Authorization"))

	var response dto.FlightOfferResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		c.JSON(http.StatusBadRequest, apierror.BadRequestErrorBindJson())
		return
	}

	c.JSON(http.StatusOK, response)
}
