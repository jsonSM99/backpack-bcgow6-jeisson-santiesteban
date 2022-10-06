package transactions

import (
	"fmt"

	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/clase06/pkg/store"
)

type Transaction struct {
	Id                int     `json:"id"`
	CodigoTransaccion string  `json:"codigoTransaccion" validate:"required"`
	Moneda            string  `json:"moneda" validate:"required"`
	Monto             float64 `json:"monto" validate:"required"`
	Emisor            string  `json:"emisor" validate:"required"`
	Receptor          string  `json:"receptor" validate:"required"`
	Fecha             string  `json:"fecha" validate:"required"`
}

var ts []Transaction

type Repository interface {
	LastId() (int, error)
	GetAll() ([]Transaction, error)
	Store(id int, codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (Transaction, error)
	Update(id int, codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (Transaction, error)
	Delete(id int) error
	Patch(id int, codigoTransaccion string, monto float64) (Transaction, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(id int, codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (Transaction, error) {
	var ts []Transaction
	r.db.Read(&ts)

	t := Transaction{
		Id:                id,
		CodigoTransaccion: codigoTransaccion,
		Moneda:            moneda,
		Emisor:            emisor,
		Receptor:          receptor,
		Fecha:             fecha,
		Monto:             monto,
	}
	ts = append(ts, t)

	if err := r.db.Write(ts); err != nil {
		return Transaction{}, err
	}
	return t, nil
}

func (r *repository) GetAll() ([]Transaction, error) {
	var ts []Transaction
	r.db.Read(&ts)
	return ts, nil
}

func (r *repository) LastId() (int, error) {
	var lastId = 0
	var ts []Transaction
	if err := r.db.Read(&ts); err != nil {
		return lastId, err
	}
	for _, t := range ts {
		if t.Id > lastId {
			lastId = t.Id
		}
	}
	return lastId + 1, nil
}

func (r *repository) Update(id int, codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (Transaction, error) {
	t := Transaction{
		CodigoTransaccion: codigoTransaccion,
		Moneda:            moneda,
		Emisor:            emisor,
		Receptor:          receptor,
		Fecha:             fecha,
		Monto:             monto,
	}

	updated := false

	for i := range ts {
		if ts[i].Id == id {
			t.Id = id
			ts[i] = t
			updated = true
		}
	}

	if !updated {
		return Transaction{}, fmt.Errorf("transacción %d no encontrada", id)
	}
	return t, nil
}

func (r *repository) Delete(id int) error {
	index := -1

	for i, t := range ts {
		if t.Id == id {
			index = i
		}
	}

	if index == -1 {
		return fmt.Errorf("transacción no encontrada")
	}

	ts = append(ts[:index], ts[index+1:]...)
	return nil
}

func (r *repository) Patch(id int, codigoTransaccion string, monto float64) (Transaction, error) {
	index := -1

	for i, t := range ts {
		if t.Id == id {
			index = i
		}
	}

	if index == -1 {
		return Transaction{}, fmt.Errorf("transacción no encontrada")
	}

	t := &ts[index]

	if len(codigoTransaccion) > 0 {
		t.CodigoTransaccion = codigoTransaccion
	}

	if monto > 0 {
		t.Monto = monto
	}
	return *t, nil
}
