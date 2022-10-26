package transactions

import (
	"testing"

	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationUpdate(t *testing.T) {
	// Arrange.
	db := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "12",
			Moneda:            "a",
			Monto:             1000,
			Emisor:            "jeisson",
			Receptor:          "fernando",
			Fecha:             "01/02/1999",
		},
	}
	updated := domain.Transaction{
		Id:                1,
		CodigoTransaccion: "UPDATED",
		Monto:             999,
	}

	mockStorage := MockStorage{
		dataMock: db,
	}

	repository := NewRepository(&mockStorage)

	service := NewService(repository)

	// Act.
	result, err := service.Patch(updated.Id, updated.CodigoTransaccion, updated.Monto)

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, true, mockStorage.ReadWasCalled, "Read no fue llamado")
	assert.Equal(t, updated.Id, result.Id, "Id: diferentes")
	assert.Equal(t, updated.CodigoTransaccion, result.CodigoTransaccion, "CodigoTransaccion: diferentes")
	assert.Equal(t, updated.Monto, result.Monto, "Monto: diferentes")
}

func TestDeleteGood(t *testing.T) {
	// Arrange.
	db := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "12",
			Moneda:            "a",
			Monto:             1000,
			Emisor:            "jeisson",
			Receptor:          "fernando",
			Fecha:             "01/02/1999",
		},
	}

	mockStorage := MockStorage{
		dataMock: db,
	}

	repository := NewRepository(&mockStorage)

	service := NewService(repository)

	// act
	errDel := service.Delete(1)

	ts, errGet := service.GetAll()
	found := false
	for _, t := range ts {
		if t.Id == 1 {
			found = true
		}
	}
	//assert
	assert.Nil(t, errDel)
	assert.Nil(t, errGet)
	assert.Equalf(t, false, found, "Transaccion no eliminada: Transaccion con Id - 1 encontrado")
	assert.Equalf(t, 0, len(ts), "Transaccion no eliminada: %v - 1 != %v", len(db), len(ts))
}
func TestDeleteBad(t *testing.T) {

}
