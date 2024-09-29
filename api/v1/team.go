package v1

type CreateTeamRequest struct {
	TeamID    string `json:"team_id" binding:"required,min=3,max=20" example:"pmr"`
	Name      string `json:"name" binding:"required,min=3,max=20" example:"PMR Fansub"`
	Desc      string `json:"desc"`
	QQGroupID string `json:"qq_group_id"`
}
