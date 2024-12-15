package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID    string `gorm:"uniqueIndex;not null"`
	Username  string `gorm:"unique;not null"`
	Nickname  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	AvatarUrl string `gorm:"not null"`
	Status    int    `gorm:"not null"`
	Teams     []Team `gorm:"many2many:user_team_role;foreignKey:UserID;joinForeignKey:UserID;References:TeamID;joinReferences:TeamID"`
}

func (u *User) TableName() string {
	return "user"
}
