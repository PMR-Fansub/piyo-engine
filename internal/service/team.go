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
	GetTeamProfile(ctx context.Context, teamID string) (*v1.TeamBasicInfo, error)
	CreateNewTeam(ctx context.Context, req *v1.CreateTeamRequest) error
	GetTeamMembers(ctx context.Context, teamID string) (*[]v1.TeamMemberProfile, error)
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

func (s *teamService) GetTeamProfile(ctx context.Context, teamID string) (*v1.TeamBasicInfo, error) {
	team, err := s.teamRepo.GetByID(ctx, teamID)
	if err != nil {
		return nil, err
	}
	return &v1.TeamBasicInfo{
		TeamID:    team.TeamID,
		Name:      team.Name,
		Status:    team.Status,
		Desc:      team.Desc,
		QQGroupID: team.QQGroupID,
		CreatedAt: team.CreatedAt,
	}, nil
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

func (s *teamService) GetTeamMembers(ctx context.Context, teamID string) (*[]v1.TeamMemberProfile, error) {
	members, err := s.teamRepo.GetMembers(ctx, teamID)
	if err != nil {
		return nil, err
	}
	var memberProfiles []v1.TeamMemberProfile
	for _, mem := range members {
		memberProfiles = append(
			memberProfiles, v1.TeamMemberProfile{
				UserBasicInfo: v1.UserBasicInfo{
					UserId:    mem.UserID,
					Username:  mem.Username,
					Nickname:  mem.Nickname,
					CreatedAt: mem.CreatedAt,
				},
				JoinedAt: mem.CreatedAt,
			},
		)
	}
	return &memberProfiles, nil
}
