package model

import "gorm.io/gorm"

type TaskService struct {
	gorm.Model
	User      UserList `gorm:"ForeignKey:Uid"`
	Uid       uint     `gorm:"not null"`
	Title     string   `gorm:"index,not null"`
	Content   string   `gorm:"type:longtext"`
	Status    int      `gorm:"default:0"`
	StartTime int64
	EndTime   int64
}
