package transactions

type Service interface {
	GetAll() ([]Transaction, error)
	Store(codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (Transaction, error)
	Update(id int, codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (Transaction, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Transaction, error) {
	ts, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (s *service) Store(codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (Transaction, error) {
	lastId, err := s.repository.LastId()

	if err != nil {
		return Transaction{}, nil
	}

	lastId++

	transaction, err := s.repository.Store(lastId, codigoTransaccion, moneda, emisor, receptor, fecha, monto)

	if err != nil {
		return Transaction{}, nil
	}

	return transaction, nil
}

func (s *service) Update(id int, codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (Transaction, error) {
	return s.repository.Update(id, codigoTransaccion, moneda, emisor, receptor, fecha, monto)
}
