package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/cmd/server/handler"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/internal/domain"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/internal/transactions"
	"github.com/jsonSM99/backpack-bcgow6-jeisson-santiesteban/go_web/integracion/test/mocks"
	"github.com/stretchr/testify/assert"
)

func createServer(mockStore mocks.MockStorage) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	repo := transactions.NewRepository(&mockStore)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	r := gin.Default()

	ts := r.Group("/transactions")
	ts.PATCH("/:id", t.Patch)
	ts.DELETE("/:id", t.Delete)

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestPatchTransaction(t *testing.T) {
	// arrrange
	t.Setenv("TOKEN", "123456")
	mockStorage := mocks.MockStorage{
		DataMock: []domain.Transaction{
			{
				Id:                1,
				CodigoTransaccion: "123",
				Moneda:            "dolar",
				Monto:             1000,
				Emisor:            "jeisson",
				Receptor:          "fernando",
				Fecha:             "01/02/1999",
			},
		},
	}
	var resp domain.Transaction

	r := createServer(mockStorage)
	req, rr := createRequestTest(http.MethodPatch, "/transactions/1", `{
		"codigoTransaccion": "12345"
    }`)
	// act
	r.ServeHTTP(rr, req)
	res := rr.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	// assert
	assert.Nil(t, err)
	err = json.Unmarshal(data, &resp)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "123", resp.CodigoTransaccion)
}

func TestDeleteTransaction(t *testing.T) {
	// arrrange
	t.Setenv("TOKEN", "123456")
	mockStorage := &mocks.MockStorage{
		DataMock: []domain.Transaction{
			{
				Id:                1,
				CodigoTransaccion: "123",
				Moneda:            "dolar",
				Monto:             1000,
				Emisor:            "jeisson",
				Receptor:          "fernando",
				Fecha:             "01/02/1999",
			},
		},
	}

	r := createServer(*mockStorage)
	req, rr := createRequestTest(http.MethodDelete, "/transactions/1", ``)
	// act
	r.ServeHTTP(rr, req)
	res := rr.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, string(data), "exito")
}
