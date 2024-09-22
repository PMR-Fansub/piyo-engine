package model

import "gorm.io/gorm"

type UserTeamRole struct {
	gorm.Model
	UserID uint   `gorm:"index"`
	TeamID uint   `gorm:"index"`
	Role   string `gorm:"not null"`
}

func (m *UserTeamRole) TableName() string {
	return "user_team_role"
}
