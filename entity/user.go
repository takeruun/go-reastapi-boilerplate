package entity

import (
	"time"
)

type User struct {
	ID           uint64 `json:"id" gorm:"primary_key"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	HashPassword string
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt"`
}
