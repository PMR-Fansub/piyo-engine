package model

import (
	"time"

	"gorm.io/gorm"
)

type UserSystemRole struct {
	UserID    uint `gorm:"primaryKey"`
	RoleID    uint `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (m *UserSystemRole) TableName() string {
	return "user_system_role"
}
