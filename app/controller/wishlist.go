package controller

import (
	"net/http"

	"github.com/felipedss/trip-wishlist-api/app/controller/apierror"
	"github.com/felipedss/trip-wishlist-api/app/controller/dto"
	"github.com/felipedss/trip-wishlist-api/domain/service"
	"github.com/gin-gonic/gin"
)

type WishlistController interface {
	Insert(c *gin.Context)
}

type WishlistControllerSt struct {
	service service.WishlistService
}

func NewWishlistController(wishlistService service.WishlistService) *WishlistControllerSt {
	return &WishlistControllerSt{
		service: wishlistService,
	}
}

func (p *WishlistControllerSt) Insert(c *gin.Context) {

	var createWishlistDTO dto.CreateWishlistDTO
	err := c.ShouldBindJSON(&createWishlistDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, apierror.BadRequestError("Error to bind json"))
		return
	}

	wishlist := createWishlistDTO.ToEntity()

	errInsert := p.service.Insert(wishlist)
	if errInsert != nil {
		c.JSON(http.StatusBadRequest, apierror.BadRequestError(errInsert.Error()))
		return
	}

	c.JSON(http.StatusOK, createWishlistDTO)
}
