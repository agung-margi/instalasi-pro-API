package invoice

type Service interface {
	GenerateInvoice(input Invoice) (Invoice, error)
	GetInvoice(id int) (Invoice, error)
	UpdateStatus(id int, status string) (Invoice, error)
	FindAll() ([]Invoice, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GenerateInvoice(input Invoice) (Invoice, error) {
	invoice := Invoice{
		OrderID:       input.OrderID,
		Total:         input.Total,
		CustomerEmail: input.CustomerEmail,
		Status:        "pending",
	}

	err := s.repository.Save(&invoice)
	if err != nil {
		return Invoice{}, err
	}
	return invoice, nil
}

func (s *service) GetInvoice(id int) (Invoice, error) {
	invoice, err := s.repository.FindById(id)
	if err != nil {
		return invoice, err
	}
	return invoice, nil
}

func (s *service) UpdateStatus(id int, status string) (Invoice, error) {
	invoice, err := s.repository.UpdateStatus(id, status)
	if err != nil {
		return invoice, err
	}
	return invoice, nil
}

func (s *service) FindAll() ([]Invoice, error) {
	invoices, err := s.repository.FindAll()
	if err != nil {
		return invoices, err
	}
	return invoices, nil
}
