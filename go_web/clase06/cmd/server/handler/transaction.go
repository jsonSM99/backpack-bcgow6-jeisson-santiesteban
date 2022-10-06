package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/clase06/internal/transactions"
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

func (c *Transaction) GetAll(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "token invalido",
		})
		return
	}

	t, err := c.service.GetAll()

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, t)
}

func (c *Transaction) Store(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "token invalido",
		})
		return
	}

	var req request
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

	t, err := c.service.Store(req.CodigoTransaccion, req.Moneda, req.Emisor, req.Receptor, req.Fecha, req.Monto)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, t)
}

func (c *Transaction) Update(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "token invalido",
		})
		return
	}

	var req request
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

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid ID"})
		return
	}
	t, err := c.service.Update(id, req.CodigoTransaccion, req.Moneda, req.Emisor, req.Receptor, req.Fecha, req.Monto)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, t)
}

func (c *Transaction) Delete(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "token invalido",
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid ID"})
		return
	}

	err = c.service.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Transacción eliminada con exito!",
	})
}

func (c *Transaction) Patch(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "token invalido",
		})
		return
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid ID"})
		return
	}
	t, err := c.service.Patch(id, req.CodigoTransaccion, req.Monto)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, t)
}
