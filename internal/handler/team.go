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

// Register godoc
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
	req := new(v1.CreateTeamRequest)
	if err := utils.BindJSONOrBadRequest(ctx, req); err != nil {
		return
	}
	req.TeamID = strings.ToLower(req.TeamID)
	if err := h.teamService.CreateNewTeam(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("teamService.CreateNewTeam error", zap.Error(err))
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}
