package models

import (
	"time"
	"gorm.io/gorm"
)

type Schedule struct {
	ID				uint		`gorm:"primaryKey" json:"id"`
	BusID			uint		`gorm:"not null" json:"bus_id"`
	RouteID			uint		`gorm:"not null" json:"route_id"`
	DepartureTime	time.Time	`gorm:"not null" json:"departure_time"`

	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		gorm.DeletedAt	`gorm:"index"`

	Bus				Bus			`gorm:"foreignKey:BusID" json:"bus"`
	Route			Route		`gorm:"foreignKey:RouteID" json:"route"`
}