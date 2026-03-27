package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID			uint		`gorm:"primaryKey" json:"id"`
	Name 		string		`gorm:"not null" json:"name"`
	Email		string		`gorm:"unique;not null" json:"email"`
	Phone		string		`gorm:"not null" json:"no_tel"`
	Password	string		`gorm:"not null" json:"-"`
	Role		string		`gorm:"type:varchar(20);default:user" json:"role"`

	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	gorm.DeletedAt `gorm:"index"`
}
