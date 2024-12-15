package v1

import (
	"time"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required_without=Email,max=30"`
	Email    string `json:"email" binding:"required_without=Username"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseData struct {
	AccessToken string `json:"access_token"`
}

type LoginResponse struct {
	Response
	Data LoginResponseData
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email" binding:"omitempty,email"`
}

type UserBasicInfo struct {
	UserId    string    `json:"user_id"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
}

type GetProfileResponseData struct {
	UserBasicInfo
}

type GetProfileResponse struct {
	Response
	Data GetProfileResponseData
}
