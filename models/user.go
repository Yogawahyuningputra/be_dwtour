package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	Phone     string    `json:"phone" gorm:"type:varchar(255)"`
	Address   string    `json:"address" gorm:"type:varchar(255)"`
	Image     string    `json:"image" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UsersResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
}

func (UsersResponse) TableName() string {
	return "users"
}