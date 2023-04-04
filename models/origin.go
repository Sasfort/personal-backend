package models

import "time"

type Origin struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	CreatedAt time.Time
}
