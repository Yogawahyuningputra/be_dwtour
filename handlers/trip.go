package handlers

import "backend/repositories"

type handlerTrip struct {
	TripRepository repositories.TripRepository
}

func HandleTrip(TripRepository repositories.TripRepository) *handlerTrip {
	return &handlerTrip{TripRepository}
}
