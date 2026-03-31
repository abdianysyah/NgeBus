package models

import (
	"time"
)

type City struct {
	ID		uint	`gorm:"primaryKey" json:"id"`
	Name	string 	`gorm:"unique;not null" json:"name"`
	Lat		float64	`gorm:"lat"`
	Lng		float64	`gorm:"lng"`

	CreatedAt time.Time
	UpdatedAt time.Time
}