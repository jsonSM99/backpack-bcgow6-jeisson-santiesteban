package mocks

import (
	"fmt"

	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/internal/domain"
)

type MockStorage struct {
	DataMock []domain.Transaction
	ErrWrite string
	ErrRead  string
}

func (m *MockStorage) Read(data interface{}) error {
	if m.ErrRead != "" {
		return fmt.Errorf(m.ErrRead)
	}
	a := data.(*[]domain.Transaction)
	*a = m.DataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.ErrWrite != "" {
		return fmt.Errorf(m.ErrWrite)
	}
	switch ts := data.(type) {
	case *[]domain.Transaction:
		m.DataMock = *ts
	default:
		return fmt.Errorf("XD")
	}
	fmt.Println(m.DataMock)
	return nil
}
