package transactions

import (
	"encoding/json"
	"testing"

	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/internal/domain"
	"github.com/stretchr/testify/assert"
)

type storeStub struct {
	data []domain.Transaction
}

func (f *storeStub) Fill() {
	f.data = []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "12",
			Moneda:            "a",
			Monto:             1000,
			Emisor:            "jeisson",
			Receptor:          "fernando",
			Fecha:             "01/02/1999",
		},
		{
			Id:                2,
			CodigoTransaccion: "122",
			Moneda:            "aaaa",
			Monto:             100011,
			Emisor:            "jeisson2",
			Receptor:          "fernando2",
			Fecha:             "02/02/1999",
		},
	}
}
func (f *storeStub) Write(data interface{}) error {
	return nil
}
func (f *storeStub) Read(data interface{}) error {
	bytes, err := json.Marshal(f.data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	return nil
}

func TestGetAll(t *testing.T) {
	// arrange
	db := storeStub{}
	db.Fill()
	repo := NewRepository(&db)

	// act
	result, err := repo.GetAll()

	//assert
	assert.Nil(t, err)
	assert.Equal(t, result, db.data, "Las transacciones %v son diferentes a %v", result, db.data)
	assert.Lenf(t, result, len(db.data), "La longitud del resultado = %v es diferente a la longitud esperada = %v", len(result), len(db.data))
}

type storeMock struct {
	StoreMockWasReaded bool
	data               []domain.Transaction
}

func (f *storeMock) Fill() {
	f.data = []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "Before Update",
			Moneda:            "a",
			Monto:             1000,
			Emisor:            "jeisson",
			Receptor:          "fernando",
			Fecha:             "01/02/1999",
		},
	}
}
func (f *storeMock) Write(data interface{}) error {
	return nil
}
func (f *storeMock) Read(data interface{}) error {
	f.StoreMockWasReaded = true
	bytes, err := json.Marshal(f.data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	return nil
}

func TestUpdate(t *testing.T) {
	// arrange
	db := storeMock{StoreMockWasReaded: false}
	db.Fill()
	repo := NewRepository(&db)

	// act
	result, err := repo.Patch(1, "AFTER UPDATE", 1999)

	//assert

	assert.Nil(t, err)
	assert.Equalf(t, true, db.StoreMockWasReaded, "Read no fue llamado")
	assert.Equalf(t, "AFTER UPDATE", result.CodigoTransaccion, "CodigoTranasaccion: no se ejecutó la actualización correctamente")
	assert.Equalf(t, float64(1999), result.Monto, "Monto: no se ejecutó la actualización correctamente")
}
