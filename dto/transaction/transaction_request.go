package transactiondto

type TransactionRequest struct {
	Qty    int    `json:"qty" gorm:"type:varchar(255)"`
	Status string `json:"status" gorm:"type:varchar(255)"`
	Image  string `json:"attachment" gorm:"type: varchar(255)"`
	Total  int    `json:"total" gorm:"type: varchar(255)"`
	TripID int    `json:"trip_id"`
	UserID int    `json:"user_id"`
}
