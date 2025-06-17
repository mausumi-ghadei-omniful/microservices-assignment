package enums

// OrderStatus represents the status of an order
type OrderStatus string

const (
	OrderStatusOnHold    OrderStatus = "on_hold"
	OrderStatusNewOrder  OrderStatus = "new_order"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

// IsValid checks if the order status is valid
func (s OrderStatus) IsValid() bool {
	switch s {
	case OrderStatusOnHold, OrderStatusNewOrder, OrderStatusCompleted, OrderStatusCancelled:
		return true
	default:
		return false
	}
}

// String returns the string representation of the order status
func (s OrderStatus) String() string {
	return string(s)
}

// InventoryOperation represents the type of inventory operation
type InventoryOperation string

const (
	InventoryOperationAdd      InventoryOperation = "add"
	InventoryOperationSubtract InventoryOperation = "subtract"
	InventoryOperationSet      InventoryOperation = "set"
)

// IsValid checks if the inventory operation is valid
func (o InventoryOperation) IsValid() bool {
	switch o {
	case InventoryOperationAdd, InventoryOperationSubtract, InventoryOperationSet:
		return true
	default:
		return false
	}
}

// String returns the string representation of the inventory operation
func (o InventoryOperation) String() string {
	return string(o)
}
