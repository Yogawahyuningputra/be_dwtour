package transactiondto

import "backend/models"

type TransactionRequest struct {
	ID         int         `json:"id"`
	Qty        int         `json:"qty" gorm:"type:varchar(255)"`
	Status     string      `json:"status" gorm:"type:varchar(255)"`
	Attachment string      `json:"image" gorm:"type: varchar(255)"`
	Total      int         `json:"total" gorm:"type: varchar(255)"`
	TripID     int         `json:"trip_id"`
	Trip       models.Trip `json:"trip"`
	UserID     int         `json:"user_id"`
}
