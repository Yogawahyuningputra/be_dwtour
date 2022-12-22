package authdto

type RegisterRequest struct {
	Fullname string `json:"fullname" gorm:"type: varchar(255)" validate:"required"`
	Email    string `json:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)" validate:"required"`
	Phone    string `json:"phone" gorm:"type: varchar(255)" validate:"required"`
	Address  string `json:"address" gorm:"type: varchar(255)" validate:"required"`
	Image    string `json:"image" gorm:"type: varchar(255)"`
}

type LoginRequest struct {
	Email    string `json:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)" validate:"required"`
}
