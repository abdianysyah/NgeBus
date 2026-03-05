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
	Price			float64		`gorm:"type:decimal(10,2);not null" json:"price"`

	CreatedAt		int64
	UpdatedAt		int64
	DeletedAt		gorm.DeletedAt	`gorm:"index"`

	Bus				Bus			`gorm:"foreignKey:BusID"`
	Route			Route		`gorm:"foreignKey:RouteID"`

	Orders			[]Order	`gorm:"foreignKey:ScheduleID"`
}