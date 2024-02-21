package order

type Service interface {
	AddOrder(customerName string, address string, amount float64) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) AddOrder(customerName string, address string, amount float64) (int64, error) {
	return s.repo.Create(customerName, address, amount)
}
