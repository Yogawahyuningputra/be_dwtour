package routes

import (
	"backend/handlers"
	"backend/pkg/middleware"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func TripRoutes(r *mux.Router) {
	tripRepository := repositories.RepositoryTrip(mysql.DB)
	h := handlers.HandlerTrip(tripRepository)
	r.HandleFunc("/trips", middleware.Auth(h.FindTrips)).Methods("GET")
	r.HandleFunc("/trip/{id}", middleware.Auth(h.GetTrip)).Methods("GET")
	r.HandleFunc("/trip", middleware.Auth(middleware.UploadFile(h.CreateTrip))).Methods("POST")
	// r.HandleFunc("/trip", middleware.Auth(h.CreateTrip)).Methods("POST")
	r.HandleFunc("/trip/{id}", middleware.Auth(h.UpdateTrip)).Methods("PATCH")
	r.HandleFunc("/trip/{id}", middleware.Auth(h.DeleteTrip)).Methods("DELETE")

}
