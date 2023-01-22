package models

import "gorm.io/gorm"

type User struct {
	ID        int32          `gorm:"primary_key:auto_increment" json:"-"`
	FitstName string         `gorm:"type:varchar(100)" json:"-"`
	LastName  string         `gorm:"type:varchar(100)" json:"-"`
	UserName  string         `gorm:"type:varchar(100);unique" json:"-"`
	Email     string         `gorm:"type:varchar(100);unique;" json:"-"`
	Password  string         `gorm:"type:varchar(100)" json:"-"`
	Delered   gorm.DeletedAt `gorm:"deleted_at"`
}
