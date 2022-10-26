package transactions

import "github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/internal/domain"

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Store(codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (domain.Transaction, error)
	Update(id int, codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (domain.Transaction, error)
	Delete(id int) error
	Patch(id int, codigoTransaccion string, monto float64) (domain.Transaction, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Transaction, error) {
	ts, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (s *service) Store(codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (domain.Transaction, error) {
	lastId, err := s.repository.LastId()

	if err != nil {
		return domain.Transaction{}, nil
	}

	lastId++

	transaction, err := s.repository.Store(lastId, codigoTransaccion, moneda, emisor, receptor, fecha, monto)

	if err != nil {
		return domain.Transaction{}, nil
	}

	return transaction, nil
}

func (s *service) Update(id int, codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (domain.Transaction, error) {
	return s.repository.Update(id, codigoTransaccion, moneda, emisor, receptor, fecha, monto)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) Patch(id int, codigoTransaccion string, monto float64) (domain.Transaction, error) {
	return s.repository.Patch(id, codigoTransaccion, monto)
}
