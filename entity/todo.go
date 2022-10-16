package entity

import "time"

type Todo struct {
	ID          uint64     `json:"id" gorm:"primary_key"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UserId      uint64     `json:"userId"`
	User        User       `json:"-"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}
