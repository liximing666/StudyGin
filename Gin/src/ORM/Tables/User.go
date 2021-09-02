package Tables

import (
	"time"
)

type User struct {
	Id uint `gorm:"primaryKey"`
	Name string `gorm:"default: test"`
	Age int `gorm:"default: 0"`
	PassWord string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	//gorm.Model
}
