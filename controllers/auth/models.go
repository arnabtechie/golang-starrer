package auth

type User struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6"`
}

type RegResp struct {
	Email string `json:"email" binding:"required,email"`
	ID    string `json:"id"`
	Token string `json:"token"`
}
