package entity

import "time"

//harus menggunakan huruf besar, karena kalau tidak dengan huruf besar maka
// ketika pemanggilan Struc user ke method lain tidak akan terpanggil
type User struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	Email     string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password  string `gorm:"->;<-;not null" json:"-"`
	Token     string `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
