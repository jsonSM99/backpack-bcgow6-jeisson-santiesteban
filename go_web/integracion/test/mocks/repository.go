package mocks

import (
	"fmt"

	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/internal/domain"
)

type MockRepository struct {
	DataMock []domain.Transaction
	Error    string
}

func (m *MockRepository) GetAll() ([]domain.Transaction, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil
}

func (m *MockRepository) Store(id int, codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (domain.Transaction, error) {
	if m.Error != "" {
		return domain.Transaction{}, fmt.Errorf(m.Error)
	}
	p := domain.Transaction{
		Id:                id,
		CodigoTransaccion: codigoTransaccion,
		Moneda:            moneda,
		Emisor:            emisor,
		Receptor:          receptor,
		Fecha:             fecha,
		Monto:             monto,
	}
	return p, nil
}

func (m *MockRepository) LastID() (int, error) {
	if m.Error != "" {
		return 0, fmt.Errorf(m.Error)
	}
	return m.DataMock[len(m.DataMock)-1].Id, nil
}
