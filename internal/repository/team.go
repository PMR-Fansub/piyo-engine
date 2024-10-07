package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
	v1 "piyo-engine/api/v1"
	"piyo-engine/internal/model"
)

type TeamRepository interface {
	Create(ctx context.Context, team *model.Team) error
	GetByID(ctx context.Context, id string) (*model.Team, error)
	GetMembers(ctx context.Context, id string) ([]model.User, error)
}

func NewTeamRepository(
	repository *Repository,
) TeamRepository {
	return &teamRepository{
		Repository: repository,
	}
}

type teamRepository struct {
	*Repository
}

func (r *teamRepository) GetByID(ctx context.Context, id string) (*model.Team, error) {
	var team model.Team
	if err := r.DB(ctx).Where("team_id = ?", id).First(&team).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &team, nil
}

func (r *teamRepository) Create(ctx context.Context, team *model.Team) error {
	if err := r.DB(ctx).Create(team).Error; err != nil {
		return err
	}
	return nil
}

func (r *teamRepository) GetMembers(ctx context.Context, id string) ([]model.User, error) {
	var team model.Team
	team.TeamID = id
	if err := r.DB(ctx).Model(&model.Team{}).Preload("Members").Find(&team).Error; err != nil {
		return nil, err
	}
	return team.Members, nil
}
