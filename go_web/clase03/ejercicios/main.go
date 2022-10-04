package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (t *Transaction) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

var transactions []Transaction

func getId() (id int) {
	for _, t := range transactions {
		if t.Id > id {
			id = t.Id
		}
	}
	id++
	return
}

func CreateTransaction(ctx *gin.Context) {

	token := ctx.GetHeader("token")

	if token != "123456" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "no tiene permisos para realizar la petición solicitada",
		})
		return
	}

	var req Transaction
	ctx.ShouldBindJSON(&req)
	if err := req.Validate(); err != nil {
		validatorErrors := err.(validator.ValidationErrors)
		r := []string{}
		for _, e := range validatorErrors {
			r = append(r, "el campo "+e.Field()+" es requerido")
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": r})
		return
	}
	req.Id = getId()
	transactions = append(transactions, req)
	ctx.JSON(http.StatusOK, req)
}

func main() {
	router := gin.Default()

	tr := router.Group("/transactions")

	tr.POST("/", CreateTransaction)

	router.Run()
}
