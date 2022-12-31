package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type TripRepository interface {
	FindTrips() ([]models.Trip, error)
	GetTrip(ID int) (models.Trip, error)
	CreateTrip(trip models.Trip) (models.Trip, error)
	UpdateTrip(trip models.Trip, ID int) (models.Trip, error)
	DeleteTrip(trip models.Trip, ID int) (models.Trip, error)
}

func RepositoryTrip(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTrips() ([]models.Trip, error) {
	var trips []models.Trip
	// err := r.db.Raw("SELECT * FROM trips").Scan(&trips).Error
	err := r.db.Preload("Country").Find(&trips).Error // ORM
	return trips, err
}

func (r *repository) GetTrip(ID int) (models.Trip, error) {
	var trip models.Trip
	// err := r.db.Raw("SELECT * FROM trips WHERE id=?", ID).Scan(&trip).Error
	err := r.db.Preload("Country").First(&trip, ID).Error // ORM
	return trip, err
}

func (r *repository) CreateTrip(trip models.Trip) (models.Trip, error) {
	// err := r.db.Exec("INSERT INTO trips(title)VALUES(?)", trip.Title).Error
	err := r.db.Preload("Country").Create(&trip).Error // ORM

	return trip, err
}
func (r *repository) UpdateTrip(trip models.Trip, ID int) (models.Trip, error) {
	err := r.db.Save(&trip).Error

	return trip, err
}
func (r *repository) DeleteTrip(trip models.Trip, ID int) (models.Trip, error) {
	err := r.db.Delete(&trip).Error
	return trip, err
}
