package routes

import (
	"backend/handlers"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func CountryRoutes(r *mux.Router) {
	countryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handlers.HandlerCountry(countryRepository)
	r.HandleFunc("/countries", h.FindCountries).Methods("GET")
	r.HandleFunc("/country/{id}", h.GetCountry).Methods("GET")
	r.HandleFunc("/country", h.CreateCountry).Methods("POST")
	r.HandleFunc("/country/{id}", h.UpdateCountry).Methods("PATCH")
	r.HandleFunc("/country/{id}", h.DeleteCountry).Methods("DELETE")

}
