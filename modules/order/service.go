package order

type Service interface {
	CreateOrder(input OrderInput) (Order, error)
	Update(id int, order Order) (Order, error)
	UpdatePickup(id int, updateOrder Order) ([]Order, error)
	UpdateProgress(id int, updateOrder Order) ([]Order, error)
	// FindAll() ([]Order, error)
	FindById(id int) (Order, error)
	// FindByUserID(id int) ([]Order, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateOrder(input OrderInput) (Order, error) {
	order := Order{
		ProductID: input.ProductID,
		UserID:    input.User.ID,
		Status:    "pending",
	}
	newOrder, err := s.repository.Create(order)
	if err != nil {
		return newOrder, err
	}
	return newOrder, nil
}

func (s *service) Update(id int, order Order) (Order, error) {
	order, err := s.repository.Update(id, order)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (s *service) FindAll() ([]Order, error) {
	orders, err := s.repository.FindAll()
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *service) FindById(id int) (Order, error) {
	order, err := s.repository.FindById(id)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (s *service) FindByUserID(id int) ([]Order, error) {
	orders, err := s.repository.FindByUserID(id)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *service) UpdatePickup(id int, updateOrder Order) ([]Order, error) {
	orders, err := s.repository.UpdatePickup(id, updateOrder)
	if err != nil {
		return orders, err
	}
	return orders, nil
}
