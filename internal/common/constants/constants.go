package constants

// Service Names
const (
	OMSServiceName = "order-management-service"
	IMSServiceName = "inventory-management-service"
)

// Order Statuses
const (
	OrderStatusOnHold    = "on_hold"
	OrderStatusNewOrder  = "new_order"
	OrderStatusCompleted = "completed"
	OrderStatusCancelled = "cancelled"
)

// Kafka Topics
const (
	TopicOrderCreated   = "order.created"
	TopicOrderFinalized = "order.finalized"
)

// SQS Queue Names
const (
	QueueCreateBulkOrder = "create-bulk-order"
)

// API Endpoints
const (
	APIVersion = "/api/v1"
)

// Database Collections/Tables
const (
	CollectionOrders = "orders"
	TableHubs        = "hubs"
	TableSKUs        = "skus"
	TableInventory   = "inventory"
)

// Cache Keys
const (
	CacheKeyHubPrefix       = "hub:"
	CacheKeySKUPrefix       = "sku:"
	CacheKeyInventoryPrefix = "inventory:"
)
