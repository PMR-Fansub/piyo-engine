package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	TeamID    string `gorm:"uniqueIndex;not null"`
	Name      string `gorm:"not null"`
	Status    int    `gorm:"not null"`
	Desc      string
	QQGroupID string
	Members   []User `gorm:"many2many:user_team_role;foreignKey:TeamID;joinForeignKey:TeamID;References:UserID;joinReferences:UserID"`
}

func (m *Team) TableName() string {
	return "team"
}
