package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID			uuid.UUIDs	`gorm:"type:char(36);primary_key" json:"id"`
	Name 		string		`gorm:"not null" json:"name"`
	Email		string		`gorm:"unique;not null" json:"email"`
	Phone		string		`gorm:"not null" json:"no_tel"`
	Password	string		`gorm:"not null" json:"-"`
	Role		string		`gorm:"type:varchar(20);default:user" json:"role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}