package customer

type Service interface {
	GetCustomers(id int) ([]Customer, error)
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

func (s *service) GetCustomers(id int) ([]Customer, error) {
	if id != 0 {
		customer, err := s.repository.GetById(id)
		if err != nil {
			return []Customer{}, err
		}
		return []Customer{customer}, nil
	}
	customers, err := s.repository.GetAll()
	if err != nil {
		return customers, err
	}
	return customers, nil
}
func (s *service) Create(customer Customer) (Customer, error) {
	customer, err := s.repository.Save(customer)
	if err != nil {
		return customer, err
	}
	return customer, nil
}
func (s *service) Update(id string, customer Customer) (Customer, error) {
	customer, err := s.repository.Update(id, customer)
	if err != nil {
		return customer, err
	}
	return customer, nil
}
func (s *service) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
