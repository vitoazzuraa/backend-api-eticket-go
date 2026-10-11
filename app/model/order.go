package model

import "time"

type Order struct {
	ID         int         `json:"id"`
	Status     string      `json:"status"`
	TotalPrice int         `json:"total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	Items      []OrderItem `json:"items,omitempty"`
}

type OrderItem struct {
	ID       int `json:"id"`
	OrderID  int `json:"order_id"`
	EventID  int `json:"event_id"`
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

type CreateOrderRequest struct {
	Items []CreateOrderItemRequest `json:"items"`
}

type CreateOrderItemRequest struct {
	EventID  int `json:"event_id"`
	Quantity int `json:"quantity"`
}
