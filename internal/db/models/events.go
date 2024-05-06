package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	Name      string    `gorm:"column:name,not null"`
	Group     Group     `gorm:"not null,foreignKey:GroupID"`
	GroupID   uint      `gorm:"column:group_id, not null"`
	Members   []User    `gorm:"many2many:event_users"`
	StartDate time.Time `gorm:"column:start_date;not null"`
	EndDate   time.Time `gorm:"column:end_date;not null"`
	Active    bool      `gorm:"column:active;default:true"`
}
