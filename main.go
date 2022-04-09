package main

import (
	"github.com/felipedss/trip-wishlist-api/infrastructure/db"
	"github.com/felipedss/trip-wishlist-api/infrastructure/routes"
	"github.com/felipedss/trip-wishlist-api/infrastructure/runtime"
	"github.com/felipedss/trip-wishlist-api/infrastructure/server"
)

func main() {
	run := runtime.InstanceRuntime()
	db.StartDB()
	s := server.NewServer()
	router := routes.ConfigRoutes(s.Server, run)
	s.Run(router)
}
