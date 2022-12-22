package models

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type: varchar(255)"`
	// User   UsersResponse `json:"user"`
	// UserID int           `json:"user_id"`
}
type CountryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}
