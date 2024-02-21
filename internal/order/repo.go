package order

import "time"

type Repository interface {
	Create(customerName string, address string, amount float64) (int64, error)
}

type repository struct {
	orders []Order
	idx    int64
}

func NewRepository() Repository {
	return &repository{idx: 1}
}

func (r *repository) Create(customerName string, address string, amount float64) (int64, error) {
	o := Order{r.idx, customerName, address, time.Now(), amount}
	r.orders = append(r.orders, o)
	r.idx++
	return o.Id, nil
}
