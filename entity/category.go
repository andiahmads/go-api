package entity

import "time"

type Categories struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Category  string `gorm:"type:varchar(255)" json:"category"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
