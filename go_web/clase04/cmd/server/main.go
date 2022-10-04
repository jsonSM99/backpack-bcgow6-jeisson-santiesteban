package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/clase04/cmd/server/handler"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/clase04/internal/transactions"
)

func main() {
	repo := transactions.NewRepository()
	service := transactions.NewService(repo)

	t := handler.NewTransaction(service)

	r := gin.Default()

	ts := r.Group("/transactions")
	ts.POST("/", t.Store)
	ts.GET("/", t.GetAll)
	ts.PUT("/:id", t.Update)

	r.Run()

}
