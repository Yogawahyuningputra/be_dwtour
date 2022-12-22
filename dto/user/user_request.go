package userdto

type CreateUserRequest struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"password" validate:"required"`
	Address  string `json:"address" form:"password" validate:"required"`
	Image    string `json:"image" form:"image"`
}

type UpdateUserRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email" `
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"password"`
	Address  string `json:"address" form:"password"`
	Image    string `json:"image" form:"image"`
}
