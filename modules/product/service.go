package product

type Service interface {
	Save(input CreateProductInput) (Product, error)
	FindAll() ([]Product, error)
	FindById(id int) (Product, error)
	UpdateStatus(id int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Save(input CreateProductInput) (Product, error) {
	product := Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		IsActive:    true,
	}
	newProduct, err := s.repository.Save(product)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (s *service) FindAll() ([]Product, error) {
	products, err := s.repository.FindAll()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *service) FindById(id int) (Product, error) {
	product, err := s.repository.FindById(id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *service) UpdateStatus(id int) (Product, error) {
	product, err := s.repository.FindById(id)
	if err != nil {
		return product, err
	}
	product.IsActive = !product.IsActive
	product, err = s.repository.Update(product.ID, product)
	if err != nil {
		return product, err
	}
	return product, nil
}
