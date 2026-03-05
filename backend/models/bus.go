package models

import "gorm.io/gorm"

type Bus struct {
	ID			uint			`gorm:"primaryKey" json:"id"`
	BusName		string			`gorm:"type:varchar(100);not null" json:"bus_name"`
	TotalSeats	int				`gorm:"not null" json:"total_seats"`

	CreatedAt	int64
	UpdatedAt	int64
	DeletedAt 	gorm.DeletedAt	`gorm:"index"`

	Schedule	[]Schedule		`gorm:"foreignKey:BusID"`
}