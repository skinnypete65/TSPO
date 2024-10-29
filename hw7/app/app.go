package app

import (
	"log"
	"net/http"

	"ecom/internal/converter"
	"ecom/internal/repository/inmemory"
	"ecom/internal/service"
	"ecom/internal/transport/rest"
	"github.com/go-playground/validator/v10"
)

const (
	serverPort = ":8080"
)

func Run() {
	goodRepo := inmemory.NewGoodRepoInMemory()
	goodService := service.NewGoodService(goodRepo)
	goodConverter := converter.GoodConverter{}
	validate := validator.New(validator.WithRequiredStructEnabled())

	goodHandler := rest.NewGoodHandler(goodService, goodConverter, validate)
	mux := newServeMux(goodHandler)

	srv := http.Server{
		Addr:    serverPort,
		Handler: mux,
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}

func newServeMux(goodHandler *rest.GoodHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /goods", goodHandler.GetAllGoods)
	mux.HandleFunc("GET /goods/{good_id}", goodHandler.GetGoodByID)
	mux.HandleFunc("POST /goods", goodHandler.AddGood)
	mux.HandleFunc("PUT /goods/{good_id}", goodHandler.UpdateGood)
	mux.HandleFunc("DELETE /goods/{good_id}", goodHandler.DeleteGoodByID)

	return mux
}
