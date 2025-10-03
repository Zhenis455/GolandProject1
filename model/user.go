package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"`
	Age       int    `gorm:"not null"`
	CreatedAt time.Time
}
