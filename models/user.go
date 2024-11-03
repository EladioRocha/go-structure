package models

import "time"

type User struct {
	ID        int       `gorm:"primaryKey"`
	Email     string
	CreatedAt time.Time
}
