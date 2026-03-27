package models

import (
	"time"
	"gorm.io/gorm"
)

type Ticket struct {
	ID				uint			`gorm:"primaryKey" json:"id"`
	OrderID			uint			`gorm:"not null" json:"order_id"`
	ScheduleID		uint			`gorm:"not null" json:"schedule_id"`
	PassengerName	string			`gorm:"type:varchar(100); not null" json:"passenger_name"`
	SeatNumber		string			`gorm:"type:varchar(100); not null" json:"seat_number"`
	Price			float64			`gorm:"type:decimal(10, 2); not null" json:"price"`

	CreatedAt 		time.Time
	UpdatedAt		time.Time
	DeletedAt		gorm.DeletedAt	`gorm:"index"`

	Order			Order			`gorm:"foreignKey:OrderID"`
	Schedule		Schedule		`gorm:"foreignKey:ScheduleID"`
}