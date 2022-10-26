package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/internal/transactions"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/pkg/web"
)

type request struct {
	CodigoTransaccion string  `json:"codigoTransaccion" validate:"required"`
	Moneda            string  `json:"moneda" validate:"required"`
	Monto             float64 `json:"monto" validate:"required"`
	Emisor            string  `json:"emisor" validate:"required"`
	Receptor          string  `json:"receptor" validate:"required"`
	Fecha             string  `json:"fecha" validate:"required"`
}

func (t *request) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

// ListTransactions godoc
// @Summary List transactions
// @Tags Products
// @Description get transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions [get]
func (c *Transaction) GetAll(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(
			http.StatusForbidden,
			web.New(
				http.StatusForbidden,
				nil,
				"token invalido",
			),
		)
		return
	}

	t, err := c.service.GetAll()

	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			web.New(
				http.StatusNotFound,
				nil,
				err.Error(),
			),
		)
		return
	}

	ctx.JSON(http.StatusOK, web.New(http.StatusOK, t, ""))
}

// StoreTransactions godoc
// @Summary Store transactions
// @Tags Transactions
// @Description store transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param transaction body request true "Transaction to store"
// @Success 200 {object} web.Response
// @Router /transactions [post]
func (c *Transaction) Store(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(
			http.StatusForbidden,
			web.New(
				http.StatusForbidden,
				nil,
				"token invalido",
			),
		)
		return
	}

	var req request
	errBinding := ctx.ShouldBindJSON(&req)
	if err := req.Validate(); errBinding != nil || err != nil {
		validatorErrors := err.(validator.ValidationErrors)
		r := []string{}
		for _, e := range validatorErrors {
			r = append(r, "el campo "+e.Field()+" es requerido")
		}
		ctx.JSON(
			http.StatusBadRequest,
			web.New(
				http.StatusBadRequest,
				nil,
				r...,
			),
		)
		return
	}

	t, err := c.service.Store(req.CodigoTransaccion, req.Moneda, req.Emisor, req.Receptor, req.Fecha, req.Monto)

	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			web.New(
				http.StatusNotFound,
				nil,
				err.Error(),
			),
		)
		return
	}

	ctx.JSON(http.StatusOK, web.New(http.StatusOK, t, ""))
}

func (c *Transaction) Update(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(
			http.StatusForbidden,
			web.New(
				http.StatusForbidden,
				nil,
				"token invalido",
			),
		)
		return
	}

	var req request
	errBinding := ctx.ShouldBindJSON(&req)
	if err := req.Validate(); errBinding != nil || err != nil {
		validatorErrors := err.(validator.ValidationErrors)
		r := []string{}
		for _, e := range validatorErrors {
			r = append(r, "el campo "+e.Field()+" es requerido")
		}
		ctx.JSON(
			http.StatusBadRequest,
			web.New(
				http.StatusBadRequest,
				nil,
				r...,
			),
		)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			web.New(
				http.StatusBadRequest,
				nil,
				"invalid ID",
			),
		)
		return
	}
	t, err := c.service.Update(id, req.CodigoTransaccion, req.Moneda, req.Emisor, req.Receptor, req.Fecha, req.Monto)

	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			web.New(
				http.StatusNotFound,
				nil,
				err.Error(),
			),
		)
		return
	}

	ctx.JSON(http.StatusOK, web.New(http.StatusOK, t, ""))
}

func (c *Transaction) Delete(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(
			http.StatusForbidden,
			web.New(
				http.StatusForbidden,
				nil,
				"token invalido",
			),
		)
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			web.New(
				http.StatusBadRequest,
				nil,
				"invalid ID",
			),
		)
		return
	}

	err = c.service.Delete(id)

	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			web.New(
				http.StatusNotFound,
				nil,
				err.Error(),
			),
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		web.New(
			http.StatusOK,
			gin.H{
				"message": "Transacción eliminada con exito!",
			},
			"",
		),
	)

}

func (c *Transaction) Patch(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	fmt.Println("TOKEN", os.Getenv("TOKEN"))
	if token != os.Getenv("TOKEN") {
		ctx.JSON(
			http.StatusForbidden,
			web.New(
				http.StatusForbidden,
				nil,
				"token invalido",
			),
		)
		return
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(
			http.StatusBadRequest,
			web.New(
				http.StatusBadRequest,
				nil,
				"invalid ID",
			),
		)
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			web.New(
				http.StatusForbidden,
				nil,
				"token invalido",
			),
		)
		return
	}
	t, err := c.service.Patch(id, req.CodigoTransaccion, req.Monto)

	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			web.New(
				http.StatusNotFound,
				nil,
				err.Error(),
			),
		)
		return
	}

	ctx.JSON(http.StatusOK, web.New(http.StatusOK, t, ""))
}
