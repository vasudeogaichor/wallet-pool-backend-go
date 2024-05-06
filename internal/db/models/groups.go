package models

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	Name        string  `gorm:"column:name;not null"`
	Members     []User  `gorm:"many2many:user_groups;"`
	Events      []Event `gorm:"one2many:group_events"`
	Active      bool    `gorm:"column:active;default:true"`
	CreatedBy   User    `gorm:"foreignkey:CreatedByID;not null"`
	CreatedByID uint    `gorm:"column:created_by_id;not null"`
}
