package authdto

type RegisterRequest struct {
	Fullname string `json:"fullname" gorm:"type: varchar(255)" validate:"required"`
	Email    string `json:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)" validate:"required"`
	Gender   string `json:"gender" gorm:"type: varchar(255)" validate:"required"`
	Phone    int    `json:"phone" gorm:"type: varchar(255)" validate:"required"`
	Address  string `json:"address" gorm:"type: varchar(255)" validate:"required"`
	Role     string `json:"role" gorm:"type: varchar(255)"`
	// Image    string `json:"image" gorm:"type: varchar(255)"`

}

type LoginRequest struct {
	Email    string `json:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)" validate:"required"`
}
