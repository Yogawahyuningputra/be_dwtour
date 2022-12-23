package transactiondto

type TransactionResponse struct {
	Qty        int    `json:"qty" gorm:"type:varchar(255)"`
	Status     string `json:"status" gorm:"type:varchar(255)"`
	Attachment string `json:"attachment" gorm:"type: varchar(255)"`
	Total      string `json:"total" gorm:"type: varchar(255)"`
	TripID     int    `json:"trip_id"`
}
