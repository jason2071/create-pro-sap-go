package models

import (
	"time"
)

type PromotionQueue struct {
	ID            uint `gorm:"primaryKey"`
	RefID         string
	GroupCode     string
	GroupNo       string
	PromotionCode string
	Filename      string
	Status        string
	InstUser      int
	CreatedAt     time.Time `gorm:"column:createdAt"`
	UpdatedAt     time.Time `gorm:"column:updatedAt"`
}

func (PromotionQueue) TableName() string {
	return "tab_promotion_queue"
}
