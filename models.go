package main

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     *string
	Age       uint8
	Birthday  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
type message struct {
	ID      uint
	Time    time.Time
	Content string
	Sender  User
	Mention *uint
}
