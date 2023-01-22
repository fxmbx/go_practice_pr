package models

type Post struct {
	ID          int32  `gorm:"primary_key:auto_increment" json:"-"`
	Title       string `gorm:"type:text;unique" json:"-"`
	Description string `gorm:"varchar(255)" json:"-"`
	UserID      int32  `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
