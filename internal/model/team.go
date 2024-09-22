package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	TeamID    int    `gorm:"unique;not null"`
	Name      string `gorm:"not null"`
	Status    int    `gorm:"not null"`
	Desc      string
	QQGroupID string
	Users     []User `gorm:"many2many:user_team_role;"`
}

func (m *Team) TableName() string {
	return "team"
}
