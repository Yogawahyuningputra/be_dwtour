package models

import "time"

type Transaction struct {
	ID         int       `json:"id"`
	Qty        int       `json:"qty" gorm:"type:varchar(255)"`
	Status     string    `json:"status" gorm:"type:varchar(255)"`
	Attachment string    `json:"attachment" gorm:"type: varchar(255)"`
	Total      string    `json:"total" gorm:"type: varchar(255)"`
	User       User      `json:"user"`
	UserID     int       `json:"user_id"`
	Trip       Trip      `json:"trip"`
	TripID     int       `json:"trip_id"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
