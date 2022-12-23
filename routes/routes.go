package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	CountryRoutes(r)
	TripRoutes(r)
	TransactionRoutes(r)

}
