package models

import "time"

type Transaction struct {
	ID         int       `json:"id"`
	Qty        int       `json:"qty" gorm:"type:varchar(255)"`
	Status     string    `json:"status" gorm:"type:varchar(255)"`
	Attachment string    `json:"image" gorm:"type: varchar(255)"`
	Total      int       `json:"total" gorm:"type: varchar(255)"`
	Trip       Trip      `json:"trip"`
	TripID     int       `json:"trip_id" gorm:"foreignKey:trip_id"`
	UserID     int       `json:"user_id" gorm:"foreignKey:user_id"`
	User       User      `json:"user" `
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
