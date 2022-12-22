package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type TripRepository interface {
	FindTrips() ([]models.Trip, error)
}

func RepositoryTrip(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTrips() ([]models.Trip, error) {
	var trips []models.Trip
	err := r.db.Raw("SELECT * FROM trips").Scan(&trips).Error
	return trips, err
}
