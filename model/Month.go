package model

import "time"

type Month struct {
	ID        uint      `json:"id"; gorm:"primary_key"`
	MonthName string    `json:"monthname"`
	CreatedAt time.Time `json:"time"`
}
