package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"column:first_name;not null"`
	LastName  string `gorm:"column:last_name"`
	Username  string `gorm:"column:username;not null;unique"`
	Email     string `gorm:"column:email;not null;unique"`
	Mobile    string `gorm:"column:mobile;unique"`
	Password  string `gorm:"column:password;not null;unique"`
}
