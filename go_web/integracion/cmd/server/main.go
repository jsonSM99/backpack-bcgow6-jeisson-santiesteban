package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/cmd/server/handler"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/docs"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/internal/transactions"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/pkg/store"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("xd")
	}
	err = godotenv.Load(filepath.Join(path, "cmd/server/.env"))

	if err != nil {
		log.Fatal(err.Error())
	}
	db := store.New(store.FileType, "./cmd/server/transacciones.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)

	t := handler.NewTransaction(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ts := r.Group("/transactions")
	ts.POST("/", t.Store)
	ts.GET("/", t.GetAll)
	ts.PUT("/:id", t.Update)
	ts.DELETE("/:id", t.Delete)
	ts.PATCH("/:id", t.Patch)

	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
