package order

import (
	"github.com/go-kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	Service
}

func NewLoggingMiddleware(logger log.Logger, s Service) Service {
	return &loggingMiddleware{logger, s}
}

func (s *loggingMiddleware) AddOrder(customerName string, address string, amount float64) (int64, error) {
	s.logger.Log("customer", customerName, "address", address, "amount", amount)
	return s.Service.AddOrder(customerName, address, amount)
}
