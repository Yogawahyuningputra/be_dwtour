package userdto

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email" `
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"password"`
	Address  string `json:"address" form:"password"`
	Image    string `json:"image" form:"image"`
}
