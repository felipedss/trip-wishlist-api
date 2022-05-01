package runtime

import (
	"github.com/felipedss/trip-wishlist-api/app/controller"
	"github.com/felipedss/trip-wishlist-api/domain/service"
	"github.com/felipedss/trip-wishlist-api/infrastructure/database/repository"
	"github.com/felipedss/trip-wishlist-api/infrastructure/external_client"
)

type Runtime struct {
	WishlistController    controller.WishlistController
	FlightOfferController controller.FlightOfferController
}

func InstanceRuntime() *Runtime {

	wishlistRepository := repository.NewWishlistRepository()
	wishlistService := service.NewWishlistService(wishlistRepository)

	flightOfferClient := external_client.NewWFlightOfferClient()
	flightOfferService := service.NewFlightOfferService(flightOfferClient)

	return &Runtime{
		WishlistController:    controller.NewWishlistController(wishlistService),
		FlightOfferController: controller.NewWFlightOfferController(flightOfferService),
	}
}
