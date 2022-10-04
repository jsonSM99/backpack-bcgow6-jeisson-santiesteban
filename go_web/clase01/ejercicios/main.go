package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// structs
type Transaction struct {
	Id                int     `json:"id"`
	CodigoTransaccion string  `json:"codigoTransaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	Fecha             string  `json:"fecha"`
}

// handlers
func MessageHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hola, Jeisson Fernando Santiesteban Mendivelso",
	})
}

func parseTransactions(path string) (data []Transaction) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		panic(err)
	}
	return data
}

func filterData(ctx *gin.Context, data []Transaction) (filtered []Transaction) {
	id := ctx.Query("id")
	codigoTransaccion := ctx.Query("codigoTransaccion")
	moneda := ctx.Query("moneda")
	monto := ctx.Query("monto")
	emisor := ctx.Query("emisor")
	receptor := ctx.Query("receptor")
	fecha := ctx.Query("fecha")
	if id == "" &&
		codigoTransaccion == "" &&
		moneda == "" &&
		monto == "" &&
		emisor == "" &&
		receptor == "" &&
		fecha == "" {
		return data
	}
	for _, t := range data {
		if _id, err := strconv.Atoi(id); err != nil && _id == t.Id {
			filtered = append(filtered, t)
			continue
		}
		if _monto, err := strconv.Atoi(monto); err != nil && _monto == int(t.Monto) {
			filtered = append(filtered, t)
			continue
		}
		if codigoTransaccion == t.CodigoTransaccion ||
			moneda == t.Moneda ||
			emisor == t.Emisor ||
			receptor == t.Receptor ||
			fecha == t.Fecha {
			filtered = append(filtered, t)
			continue
		}
	}
	return
}

func GetAll(ctx *gin.Context) {
	data := parseTransactions("transacciones.json")

	filteredData := filterData(ctx, data)

	ctx.JSON(200, gin.H{
		"accepted": filteredData,
	})
}

func GetOne(ctx *gin.Context) {
	data := parseTransactions("transacciones.json")

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": id,
		})
		return
	}
	var transaction Transaction
	for _, t := range data {
		if _id, err := strconv.Atoi(id); err == nil && t.Id == _id {
			transaction = t
		}
	}
	ctx.JSON(200, gin.H{
		"accepted": transaction,
	})
}

// main
func main() {
	router := gin.Default()

	router.GET("/greetings", MessageHandler)
	router.GET("/transactions", GetAll)
	router.GET("/transactions/:id", GetOne)

	router.Run()
}
