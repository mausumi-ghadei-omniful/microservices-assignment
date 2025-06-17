package models

import (
	"time"

	"gorm.io/gorm"
)

// Hub represents a distribution hub
type Hub struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Code       string         `json:"code" gorm:"uniqueIndex;not null"`
	Name       string         `json:"name" gorm:"not null"`
	Address    string         `json:"address"`
	City       string         `json:"city"`
	State      string         `json:"state"`
	Country    string         `json:"country"`
	PostalCode string         `json:"postal_code"`
	Phone      string         `json:"phone"`
	Email      string         `json:"email"`
	IsActive   bool           `json:"is_active" gorm:"default:true"`
	TenantID   string         `json:"tenant_id" gorm:"index"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName specifies the table name for Hub
func (Hub) TableName() string {
	return "hubs"
}
