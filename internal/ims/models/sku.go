package models

import (
	"time"

	"gorm.io/gorm"
)

// SKU represents a Stock Keeping Unit
type SKU struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Code        string         `json:"code" gorm:"uniqueIndex;not null"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Brand       string         `json:"brand"`
	Weight      float64        `json:"weight"`
	Dimensions  string         `json:"dimensions"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	TenantID    string         `json:"tenant_id" gorm:"index"`
	SellerID    string         `json:"seller_id" gorm:"index"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName specifies the table name for SKU
func (SKU) TableName() string {
	return "skus"
}
