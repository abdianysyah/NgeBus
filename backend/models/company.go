package models

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID				uint			`gorm:"primaryKey" json:"id"`
	NameCompany		string			`gorm:"type:varchar(80) ;not null" json:"name_company"`
	// TipeCompany		string			`gorm:"type:enum('AKAP', 'AKDP', 'AJDP'); not null" json:"tipe_company"`
	TotalBus		int				`gorm:"not null" json:"total_bus"`

	CreatedAt 		time.Time 
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt 	`gorm:"index"`

	Bus				[]Bus			`gorm:"foreignKey:CompanyID"`
}