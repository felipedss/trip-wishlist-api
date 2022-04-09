package routes

import (
	"github.com/felipedss/trip-wishlist-api/infrastructure/runtime"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, run *runtime.Runtime) *gin.Engine {

	main := router.Group("/v1")
	{
		wishlist := main.Group("wishlist")
		{
			wishlist.POST("/", run.WishlistController.Insert)
		}
	}
	return router
}
