package tripdto

type TripResponse struct {
	ID             int    `json:"id"`
	Title          string `json:"title" gorm:"type:varchar(255)"`
	CountryID      int    `json:"country_id" `
	Acomodation    string `json:"acomodation" gorm:"type:varchar(255)"`
	Transportation string `json:"transportation" gorm:"type:varchar(255)"`
	Eat            string `json:"eat" gorm:"type:varchar(255)"`
	Day            int    `json:"day" gorm:"type:varchar(255)"`
	Night          int    `json:"night" gorm:"type:varchar(255)"`
	DateTrip       string `json:"date_trip" gorm:"type:varchar(255)"`
	Price          int    `json:"price" gorm:"type:varchar(255)"`
	Quota          int    `json:"quota" gorm:"type:varchar(255)"`
	Description    string `json:"description" gorm:"type:varchar(255)"`
	Image          string `json:"image" form:"image" gorm:"type:varchar(255)"`
}
