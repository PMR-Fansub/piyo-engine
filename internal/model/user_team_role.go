package model

import (
	"time"

	"gorm.io/gorm"
)

type UserTeamRole struct {
	UserID    uint `gorm:"primaryKey"`
	TeamID    uint `gorm:"primaryKey"`
	Role      int  `gorm:"not null"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (m *UserTeamRole) TableName() string {
	return "user_team_role"
}
