package mocks

import (
	"fmt"

	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/internal/domain"
)

type MockService struct {
	DataMock []domain.Transaction
	Error    string
}

func (m *MockService) GetAll() ([]domain.Transaction, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil

}

func (m *MockService) Store(codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (domain.Transaction, error) {
	if m.Error != "" {
		return domain.Transaction{}, fmt.Errorf(m.Error)
	}
	t := domain.Transaction{
		CodigoTransaccion: codigoTransaccion,
		Moneda:            moneda,
		Emisor:            emisor,
		Receptor:          receptor,
		Fecha:             fecha,
		Monto:             monto,
	}
	return t, nil
}
