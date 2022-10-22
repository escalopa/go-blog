package entities

import "time"

type Post struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	CreateAt    time.Time `json:"create-at"`
	UpdatedAt   time.Time `json:"updated-at"`
}
