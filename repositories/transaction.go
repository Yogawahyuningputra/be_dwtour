package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	UpdateStatus(status string, ID int) error
	// GetOneTransaction(ID string) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.Country").Preload("User").Find(&transactions).Error
	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.Country").Preload("User").First(&transaction, ID).Error
	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}
func (r *repository) UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Model(&transaction).Updates(transaction).Error
	// err := r.db.Raw(`UPDATE transactions SET qty = ?, status = ?, attachment = ?, total = ?, trip_id = ?, user_id = ?, created_at = ?, updated_at = ? WHERE transactions.id = ?`, transaction.Qty, transaction.Status, transaction.Attachment, transaction.Total, transaction.TripID, transaction.UserID, transaction.CreatedAt, transaction.UpdatedAt, transaction.ID).Scan(&transaction).Error
	return transaction, err
}
func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error
	return transaction, err
}

func (r *repository) UpdateStatus(status string, ID int) error {
	var transaction models.Transaction
	r.db.Preload("Trip").First(&transaction, ID)

	if status != transaction.Status && status == "success" {
		var trip models.Trip
		r.db.First(&trip, transaction.Trip.ID)
		trip.Quota = trip.Quota - transaction.Qty
		r.db.Save(&trip)
	}
	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}

// func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
// 	var transaction models.Transaction
// 	err := r.db.Preload("Trip").Preload("Trip.Country").Preload("User").First(&transaction, "id = ?", ID).Error

// 	return transaction, err
// }
