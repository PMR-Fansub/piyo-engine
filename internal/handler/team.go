package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"piyo-engine/api/v1"
	"piyo-engine/internal/service"
	"piyo-engine/internal/utils"
)

type TeamHandler struct {
	*Handler
	teamService service.TeamService
}

func NewTeamHandler(
	handler *Handler,
	teamService service.TeamService,
) *TeamHandler {
	return &TeamHandler{
		Handler:     handler,
		teamService: teamService,
	}
}

// CreateTeam godoc
// @Summary 新建团队
// @Schemes
// @Description
// @Tags 团队模块
// @Accept json
// @Produce json
// @Param request body v1.CreateTeamRequest true "params"
// @Success 200 {object} v1.Response
// @Router /team [post]
func (h *TeamHandler) CreateTeam(ctx *gin.Context) {
	var req v1.CreateTeamRequest
	if err := utils.BindJSONOrBadRequest(ctx, &req); err != nil {
		return
	}
	req.TeamID = strings.ToLower(req.TeamID)
	if err := h.teamService.CreateNewTeam(ctx, &req); err != nil {
		h.logger.WithContext(ctx).Error("teamService.CreateNewTeam error", zap.Error(err))
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

// GetTeamProfile godoc
// @Summary 获取团队信息
// @Schemes
// @Description
// @Tags 团队模块
// @Accept json
// @Produce json
// @Param team_id path string true "Team ID"
// @Success 200 {object} v1.GetTeamProfileResponseData
// @Router /team/{team_id} [get]
func (h *TeamHandler) GetTeamProfile(ctx *gin.Context) {
	teamID := ctx.Param("team_id")
	teamProfile, err := h.teamService.GetTeamProfile(ctx, strings.ToLower(teamID))
	if err != nil {
		h.logger.WithContext(ctx).Error("teamService.GetTeamProfile error", zap.Error(err))
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(
		ctx, v1.GetTeamProfileResponseData{
			TeamBasicInfo: *teamProfile,
		},
	)
}

// GetTeamMembers godoc
// @Summary 获取团队成员列表
// @Schemes
// @Description
// @Tags 团队模块
// @Accept json
// @Produce json
// @Param team_id path string true "Team ID"
// @Success 200 {object} v1.GetTeamMembersResponseData
// @Router /team/{team_id}/members [get]
func (h *TeamHandler) GetTeamMembers(ctx *gin.Context) {
	teamID := ctx.Param("team_id")
	members, err := h.teamService.GetTeamMembers(ctx, strings.ToLower(teamID))
	if err != nil {
		h.logger.WithContext(ctx).Error("teamService.GetTeamMembers error", zap.Error(err))
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(
		ctx, v1.GetTeamMembersResponseData{
			Members: *members,
		},
	)
}
