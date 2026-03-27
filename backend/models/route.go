package models

import (
	"time"

	"gorm.io/gorm"
)

type Route struct {
	ID 				uint			`gorm:"primaryKey" json:"id"`
	Origin			string			`gorm:"type:varchar(100);not null" json:"origin"`
	Destination		string			`gorm:"type:varchar(100);not null" json:"destination"`
	Distance		int				`gorm:"not null" json:"distance"`

	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt 		gorm.DeletedAt	`gorm:"index"`

	Schedule		[]Schedule		`gorm:"foreignKey:RouteID"`
}