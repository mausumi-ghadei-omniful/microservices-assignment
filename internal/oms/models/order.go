package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order represents an order in the system
type Order struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OrderID     string             `json:"order_id" bson:"order_id,omitempty"`
	SKUCode     string             `json:"sku_code" bson:"sku_code"`
	HubCode     string             `json:"hub_code" bson:"hub_code"`
	Quantity    int                `json:"quantity" bson:"quantity"`
	SellerID    string             `json:"seller_id" bson:"seller_id"`
	TenantID    string             `json:"tenant_id" bson:"tenant_id"`
	Status      string             `json:"status" bson:"status"`
	CustomerID  string             `json:"customer_id" bson:"customer_id,omitempty"`
	TotalAmount float64            `json:"total_amount" bson:"total_amount,omitempty"`
	Currency    string             `json:"currency" bson:"currency,omitempty"`
	Notes       string             `json:"notes" bson:"notes,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

// OrderValidation represents validation results for orders
type OrderValidation struct {
	OrderID     string `json:"order_id" bson:"order_id"`
	SKUCode     string `json:"sku_code" bson:"sku_code"`
	HubCode     string `json:"hub_code" bson:"hub_code"`
	Quantity    int    `json:"quantity" bson:"quantity"`
	IsValid     bool   `json:"is_valid" bson:"is_valid"`
	ErrorReason string `json:"error_reason" bson:"error_reason,omitempty"`
}

// BulkOrderRequest represents a bulk order upload request
type BulkOrderRequest struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	JobID       string             `json:"job_id" bson:"job_id"`
	S3FilePath  string             `json:"s3_file_path" bson:"s3_file_path"`
	Status      string             `json:"status" bson:"status"`
	TotalRows   int                `json:"total_rows" bson:"total_rows"`
	ValidRows   int                `json:"valid_rows" bson:"valid_rows"`
	InvalidRows int                `json:"invalid_rows" bson:"invalid_rows"`
	TenantID    string             `json:"tenant_id" bson:"tenant_id"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	CompletedAt *time.Time         `json:"completed_at" bson:"completed_at,omitempty"`
}

// Webhook represents webhook registration
type Webhook struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	URL       string             `json:"url" bson:"url"`
	Events    []string           `json:"events" bson:"events"`
	TenantID  string             `json:"tenant_id" bson:"tenant_id"`
	IsActive  bool               `json:"is_active" bson:"is_active"`
	Secret    string             `json:"secret" bson:"secret"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
