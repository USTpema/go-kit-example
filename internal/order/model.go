package order

import "time"

type Order struct {
	Id           int64     `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	Address      string    `json:"address"`
	DateOfOrder  time.Time `json:"date_of_order"`
	Amount       float64   `json:"amount"`
}
