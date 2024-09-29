package service

import (
	"context"
	"errors"

	v1 "piyo-engine/api/v1"
	"piyo-engine/internal/constant"
	"piyo-engine/internal/model"
	"piyo-engine/internal/repository"
)

type TeamService interface {
	GetTeam(ctx context.Context, id string) (*model.Team, error)
	CreateNewTeam(ctx context.Context, req *v1.CreateTeamRequest) error
}

func NewTeamService(
	service *Service,
	teamRepository repository.TeamRepository,
) TeamService {
	return &teamService{
		Service:  service,
		teamRepo: teamRepository,
	}
}

type teamService struct {
	*Service
	teamRepo repository.TeamRepository
}

func (s *teamService) GetTeam(ctx context.Context, id string) (*model.Team, error) {
	return s.teamRepo.GetByID(ctx, id)
}

func (s *teamService) CreateNewTeam(ctx context.Context, req *v1.CreateTeamRequest) error {
	team, err := s.teamRepo.GetByID(ctx, req.TeamID)
	if err != nil && !errors.Is(err, v1.ErrNotFound) {
		return v1.ErrInternalServerError
	}
	if team != nil {
		return v1.ErrTeamIDAlreadyUse
	}
	team = &model.Team{
		TeamID:    req.TeamID,
		Name:      req.Name,
		Status:    constant.TeamStatusNormal,
		Desc:      req.Desc,
		QQGroupID: req.QQGroupID,
	}
	err = s.tm.Transaction(
		ctx, func(ctx context.Context) error {
			if err = s.teamRepo.Create(ctx, team); err != nil {
				return err
			}
			return nil
		},
	)
	return err
}
