package userdto

type UserResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email" `
	Password string `json:"password" form:"password"`
	Gender   string `json:"gender" form:"gender"`
	Phone    int    `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Image    string `json:"image" form:"image"`
	Role     string `json:"role" form:"role"`
}
