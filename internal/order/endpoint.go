package order

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addOrderRequest struct {
	CustomerName string
	Address      string
	Amount       float64
}

type addOrderResponse struct {
	Id int64 `json:"order_id"`
}

func makeAddOrderEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(addOrderRequest)
		id, err := s.AddOrder(req.CustomerName, req.Address, req.Amount)
		return addOrderResponse{id}, err
	}
}
