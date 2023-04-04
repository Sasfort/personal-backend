package models

import "time"

type Character struct {
	Id         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Appearance string `json:"appearance"`
	OriginId   uint   `json:"origin_id"`
	Origin     Origin `gorm:"foreignKey:OriginId"`
	CreatedAt  time.Time
}
