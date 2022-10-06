package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/clase07/cmd/server/handler"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/clase07/internal/transactions"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/clase07/pkg/store"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}
	db := store.New(store.FileType, "./transacciones.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)

	t := handler.NewTransaction(service)

	r := gin.Default()

	ts := r.Group("/transactions")
	ts.POST("/", t.Store)
	ts.GET("/", t.GetAll)
	ts.PUT("/:id", t.Update)
	ts.DELETE("/:id", t.Delete)
	ts.PATCH("/:id", t.Patch)

	r.Run()

}
