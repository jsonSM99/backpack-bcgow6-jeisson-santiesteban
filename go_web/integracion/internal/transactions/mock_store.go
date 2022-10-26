package transactions

import (
	"encoding/json"

	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/internal/domain"
)

type MockStorage struct {
	dataMock       []domain.Transaction
	ReadWasCalled  bool
	WriteWasCalled bool
	errOnWrite     error
	errOnRead      error
}

func (m *MockStorage) Read(data interface{}) (err error) {
	m.ReadWasCalled = true
	if m.errOnRead != nil {
		return m.errOnRead
	}

	bytedata, err := json.Marshal(m.dataMock)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytedata, &data)

	if err != nil {
		return err
	}
	return nil
}

func (m *MockStorage) Write(data interface{}) (err error) {
	m.WriteWasCalled = true
	if m.errOnWrite != nil {
		return m.errOnWrite
	}

	dataBytes, err := json.Marshal(data)

	if err != nil {
		return err
	}

	var castedData []domain.Transaction
	err = json.Unmarshal(dataBytes, &castedData)

	if err != nil {
		return err
	}

	m.dataMock = castedData

	return nil
}
