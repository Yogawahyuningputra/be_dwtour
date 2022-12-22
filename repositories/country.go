package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type CountryRepository interface {
	FindCountries() ([]models.Country, error)
	GetCountry(ID int) (models.Country, error)
	CreateCountry(country models.Country) (models.Country, error)
	UpdateCountry(country models.Country, ID int) (models.Country, error)
	DeleteCountry(country models.Country, ID int) (models.Country, error)
}

func RepositoryCountry(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCountries() ([]models.Country, error) {
	var countries []models.Country
	// err := r.db.Raw("SELECT * FROM countries").Scan(&countries).Error
	err := r.db.Find(&countries).Error // with ORM
	return countries, err
}
func (r *repository) GetCountry(ID int) (models.Country, error) {
	var country models.Country
	// err := r.db.Raw("SELECT * FROM countries WHERE id=?", ID).Scan(&country).Error
	err := r.db.First(&country, ID).Error
	return country, err
}
func (r *repository) CreateCountry(country models.Country) (models.Country, error) {
	// err := r.db.Exec("INSERT INTO countries(name)VALUES(?)", country.Name).Error
	err := r.db.Create(&country).Error // ORM

	return country, err
}

func (r *repository) UpdateCountry(country models.Country, ID int) (models.Country, error) {
	// err := r.db.Raw("UPDATE countries SET name=? WHERE id=?", country.Name, ID).Scan(&country).Error
	err := r.db.Save(&country).Error
	return country, err
}

func (r *repository) DeleteCountry(country models.Country, ID int) (models.Country, error) {
	err := r.db.Raw("DELETE  FROM countries WHERE id=?", ID).Scan(&country).Error
	// err := r.db.Delete(&country).Error
	return country, err
}
