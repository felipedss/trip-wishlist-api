package controller

import (
	"encoding/json"
	"github.com/felipedss/trip-wishlist-api/app/controller/apierror"
	"github.com/felipedss/trip-wishlist-api/app/controller/dto"
	"github.com/felipedss/trip-wishlist-api/app/controller/rest"
	"github.com/felipedss/trip-wishlist-api/infrastructure/env"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type FlightOfferController interface {
	GetAll(c *gin.Context)
}

type FlightOfferControllerSt struct {
}

func NewWFlightOfferController() *WishlistControllerSt {
	return &WishlistControllerSt{}
}

func (p *WishlistControllerSt) GetAll(c *gin.Context) {

	urlPath, err := rest.BuildQueryParam(env.GetEnvConfig().Url, c.Request.URL.Query())
	if err != nil {
		c.JSON(http.StatusBadRequest, apierror.BadRequestErrorQueryParms())
		return
	}

	req, _ := http.NewRequest("GET", urlPath, nil)

	req.Header.Add("accept", c.GetHeader("accept"))
	req.Header.Add("Authorization", c.GetHeader("Authorization"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, apierror.BadRequestErrorCallRemote())
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var response dto.FlightOfferResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		c.JSON(http.StatusBadRequest, apierror.BadRequestErrorBindJson())
		return
	}

	c.JSON(http.StatusOK, response)
}
