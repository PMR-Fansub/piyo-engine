package v1

import (
	"time"
)

type CreateTeamRequest struct {
	TeamID    string `json:"team_id" binding:"required,min=3,max=20" example:"pmr"`
	Name      string `json:"name" binding:"required,min=3,max=20" example:"PMR Fansub"`
	Desc      string `json:"desc"`
	QQGroupID string `json:"qq_group_id"`
}

type GetTeamProfileResponseData struct {
	TeamID    string    `json:"team_id"`
	Name      string    `json:"name"`
	Status    int       `json:"status"`
	Desc      string    `json:"desc"`
	QQGroupID string    `json:"qq_group_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetTeamProfileResponse struct {
	Response
	Data GetProfileResponseData
}

type TeamMemberProfile struct {
	GetProfileResponseData
	JoinedAt time.Time `json:"joined_at"`
}

type GetTeamMembersResponseData struct {
	Members []TeamMemberProfile `json:"members"`
}

type GetTeamMembersResponse struct {
	Response
	Data GetTeamMembersResponseData
}
