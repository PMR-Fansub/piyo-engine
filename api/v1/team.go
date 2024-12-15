package v1

import (
	"time"
)

type CreateTeamRequest struct {
	TeamID    string `json:"team_id" binding:"required,min=3,max=20"`
	Name      string `json:"name" binding:"required,min=3,max=20"`
	Desc      string `json:"desc"`
	QQGroupID string `json:"qq_group_id"`
}

type TeamBasicInfo struct {
	TeamID    string    `json:"team_id"`
	Name      string    `json:"name"`
	Status    int       `json:"status"`
	Desc      string    `json:"desc"`
	QQGroupID string    `json:"qq_group_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetTeamProfileResponseData struct {
	TeamBasicInfo
}

type GetTeamProfileResponse struct {
	Response
	Data GetTeamProfileResponseData
}

type TeamMemberProfile struct {
	UserBasicInfo
	JoinedAt time.Time `json:"joined_at"`
}

type GetTeamMembersResponseData struct {
	Members []TeamMemberProfile `json:"members"`
}

type GetTeamMembersResponse struct {
	Response
	Data GetTeamMembersResponseData
}
