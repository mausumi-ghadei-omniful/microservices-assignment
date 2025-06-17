package models

import (
	"time"

	"gorm.io/gorm"
)

// Inventory represents inventory levels for SKUs at specific hubs
type Inventory struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	HubID        uint           `json:"hub_id" gorm:"not null;index"`
	SKUID        uint           `json:"sku_id" gorm:"not null;index"`
	Quantity     int            `json:"quantity" gorm:"default:0"`
	Reserved     int            `json:"reserved" gorm:"default:0"`
	Available    int            `json:"available" gorm:"default:0"`
	MinThreshold int            `json:"min_threshold" gorm:"default:0"`
	MaxThreshold int            `json:"max_threshold"`
	TenantID     string         `json:"tenant_id" gorm:"index"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// Relationships
	Hub Hub `json:"hub,omitempty" gorm:"foreignKey:HubID"`
	SKU SKU `json:"sku,omitempty" gorm:"foreignKey:SKUID"`
}

// TableName specifies the table name for Inventory
func (Inventory) TableName() string {
	return "inventory"
}

// BeforeSave calculates available quantity
func (i *Inventory) BeforeSave(tx *gorm.DB) error {
	i.Available = i.Quantity - i.Reserved
	return nil
}
