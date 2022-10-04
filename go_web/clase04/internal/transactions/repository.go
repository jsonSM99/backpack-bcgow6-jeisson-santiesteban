package transactions

import "fmt"

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
var lastId int

type Repository interface {
	LastId() (int, error)
	GetAll() ([]Transaction, error)
	Store(id int, codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (Transaction, error)
	Update(id int, codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (Transaction, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(id int, codigoTransaccion, moneda, emisor, receptor, fecha string, monto float64) (Transaction, error) {
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
	lastId = t.Id
	return t, nil
}

func (r *repository) GetAll() ([]Transaction, error) {
	return ts, nil
}

func (r *repository) LastId() (int, error) {
	return lastId, nil
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
		return Transaction{}, fmt.Errorf("Transaccion %d no encontrada", id)
	}
	return t, nil
}
