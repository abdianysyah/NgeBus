package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID			uuid.UUID	`gorm:"type:char(36);primary_key" json:"id"`
	Name 		string		`gorm:"not null" json:"name"`
	Email		string		`gorm:"unique;not null" json:"email"`
	Phone		string		`gorm:"not null" json:"no_tel"`
	Password	string		`gorm:"not null" json:"-"`
	Role		string		`gorm:"type:varchar(20);default:user" json:"role"`
}
