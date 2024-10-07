package v1

import "time"

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=30" example:"foo"`
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required_without=Email,max=30" example:"foo"`
	Email    string `json:"email" binding:"required_without=Username" example:"1234@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LoginResponseData struct {
	AccessToken string `json:"access_token"`
}

type LoginResponse struct {
	Response
	Data LoginResponseData
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname" example:"alan"`
	Email    string `json:"email" binding:"omitempty,email" example:"1234@gmail.com"`
}

type GetProfileResponseData struct {
	UserId    string    `json:"user_id"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname" example:"alan"`
	CreatedAt time.Time `json:"created_at"`
}

type GetProfileResponse struct {
	Response
	Data GetProfileResponseData
}
