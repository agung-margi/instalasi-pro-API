package customer

type Service interface {
	GetAll() ([]Customer, error)
	GetById(id int) (Customer, error)
	Create(customer Customer) (Customer, error)
	Update(id string, customer Customer) (Customer, error)
	Delete(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Customer, error) {
	customers, err := s.repository.GetAll()
	return customers, err
}

func (s *service) GetById(id int) (Customer, error) {
	customer, err := s.repository.GetById(id)
	return customer, err
}
