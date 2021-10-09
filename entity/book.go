package entity

import "time"

type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	//buat relasi ketable user
	UserID string `gorm:"not null" json:"-"`
	//definisikan struct users
	User      User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
