package models

import (
	"time"
	"gorm.io/gorm"
)

type Order struct {
	ID         	uint           `gorm:"primaryKey" json:"id"`
	UserID     	uint           `gorm:"not null" json:"user_id"`
	ScheduleID 	uint           `gorm:"not null;uniqueIndex:seat_booking" json:"schedule_id"`
	SeatNumber 	int            `gorm:"not null;uniqueIndex:seat_booking" json:"seat_number"`
	Status     	string         `gorm:"type:varchar(20);default:'pending'" json:"status"`
	
	CreatedAt  	time.Time
	UpdatedAt  	time.Time
	DeletedAt  	gorm.DeletedAt `gorm:"index"`
	
	User       	User           `gorm:"foreignKey:UserID"`
	Schedule   	Schedule       `gorm:"foreignKey:ScheduleID"`
}