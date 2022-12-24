package tripdto

type TripRequest struct {
	ID             int    `json:"id"`
	Title          string `json:"title" form:"title" gorm:"type:varchar(255)"`
	CountryID      int    `json:"country_id"`
	Acomodation    string `json:"acomodation" form:"acomodation" gorm:"type:varchar(255)"`
	Transportation string `json:"transportation" form:"transportation" gorm:"type:varchar(255)"`
	Eat            string `json:"eat" form:"eat" gorm:"type:varchar(255)"`
	Day            string `json:"day" form:"day" gorm:"type:varchar(255)"`
	Night          string `json:"night" form:"night" gorm:"type:varchar(255)"`
	DateTrip       string `json:"date_trip" form:"date_trip" gorm:"type:varchar(255)"`
	Price          int    `json:"price" form:"price" gorm:"type:varchar(255)"`
	Quota          int    `json:"quota" form:"quota" gorm:"type:varchar(255)"`
	Description    string `json:"description" form:"description" gorm:"type:varchar(255)"`
	Image          string `json:"image" form:"image" gorm:"type:varchar(255)" `
}
