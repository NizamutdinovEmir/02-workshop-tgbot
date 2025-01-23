package product

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	// ToDo проверить на ошибки
	return &allProducts[idx], nil
}
