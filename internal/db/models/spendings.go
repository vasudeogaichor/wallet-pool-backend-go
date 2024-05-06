package models

import (
	"github.com/jinzhu/gorm"
)

type Spending struct {
	gorm.Model
	Description string  `gorm:"column:description,not null"`
	Amount      float64 `gorm:"column:amount;not null,numeric"`
	GroupID     uint    `gorm:"column:group_id;not null"`
	EventID     uint    `gorm:"column:event_id;not null"`
	Status      string  `gorm:"column:status;not null"`
	CreatedBy   User    `gorm:"foreignkey:CreatedByID"`
	CreatedByID uint    `gorm:"column:created_by_id;not null"`
}
