package repository

import (
	"context"

	"piyo-engine/internal/model"
)

type UserTeamRoleRepository interface {
	GetUserTeamRole(ctx context.Context, id int64) (*model.UserTeamRole, error)
}

func NewUserTeamRoleRepository(
	repository *Repository,
) UserTeamRoleRepository {
	return &userTeamRoleRepository{
		Repository: repository,
	}
}

type userTeamRoleRepository struct {
	*Repository
}

func (r *userTeamRoleRepository) GetUserTeamRole(ctx context.Context, id int64) (*model.UserTeamRole, error) {
	var userTeamRole model.UserTeamRole
	return &userTeamRole, nil
}
