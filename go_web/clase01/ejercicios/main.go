package main

import (
	"encoding/json"
	"io/ioutil"

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

func GetAll(ctx *gin.Context) {
	data := parseTransactions("transacciones.json")

	ctx.JSON(200, gin.H{
		"accepted": data,
	})
}

// main
func main() {
	router := gin.Default()

	router.GET("/greetings", MessageHandler)
	router.GET("/transactions", GetAll)

	router.Run()
}
