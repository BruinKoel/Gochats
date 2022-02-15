package code

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Age       *uint8
	Birthday  *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Message struct {
	ID      uint `gorm:"primaryKey"`
	Time    *time.Time
	Content string
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
	Mention *uint
}
