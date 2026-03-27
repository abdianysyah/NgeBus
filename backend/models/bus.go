package models

import (
	"time"

	"gorm.io/gorm"
)

type Bus struct {
	ID			uint			`gorm:"primaryKey" json:"id"`
	CompanyID	uint			`gorm:"not null" json:"company_id"`
	BusName		string			`gorm:"type:varchar(100);not null" json:"bus_name"`
	BusNumber	string			`gorm:"type:varchar(50);unique" json:"bus_number"`
	BusClass	string			`gorm:"type:enum('Ekonomi', 'Bisnis', 'Executive'); not null" json:"bus_class"`
	Status		string			`gorm:"type:varchar(20);default:'active'" json:"status"`
	TotalSeats	int				`gorm:"not null" json:"total_seats"`

	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt 	gorm.DeletedAt	`gorm:"index"`

	Schedule	[]Schedule		`gorm:"foreignKey:BusID"`
}