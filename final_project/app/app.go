package app

import (
	"log"
	"net/http"

	"ecom/internal/converter"
	"ecom/internal/repository/postgresdb"
	"ecom/internal/service"
	"ecom/internal/tokens"
	"ecom/internal/transport/rest"
	"ecom/internal/transport/rest/middleware"
	"ecom/pkg/hash"
	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	serverPort = ":8080"
)

func Run() {
	db, err := createGormDB()
	if err != nil {
		log.Fatal(err)
	}

	goodRepo := postgresdb.NewGoodPostgresRepo(db)
	goodService := service.NewGoodService(goodRepo)
	goodConverter := converter.GoodConverter{}
	validate := validator.New(validator.WithRequiredStructEnabled())

	hasher := hash.NewSHA1Hasher("someSalt")
	tokenManager, err := tokens.NewTokenManager("someSigningKey")
	if err != nil {
		log.Fatal(err)
	}

	authRepo := postgresdb.NewAuthRepoPostgres(db)
	authService := service.NewAuthService(authRepo, hasher, tokenManager)

	paginationRepo := postgresdb.NewPaginationRepoPostgres(db)
	paginationService := service.NewPaginationService(paginationRepo)

	goodHandler := rest.NewGoodHandler(goodService, paginationService, goodConverter, validate)
	authHandler := rest.NewAuthHandler(authService, validate)
	authMiddleware := middleware.NewAuthMiddleware(tokenManager)

	mux := newServeMux(goodHandler, authHandler, authMiddleware)

	srv := http.Server{
		Addr:    serverPort,
		Handler: mux,
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func newServeMux(
	goodHandler *rest.GoodHandler,
	authHandler *rest.AuthHandler,
	authMiddleware *middleware.AuthMiddleware,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /goods", authMiddleware.CheckAuth(
		http.HandlerFunc(goodHandler.GetAllGoods),
	))
	mux.Handle("GET /goods/{good_id}", authMiddleware.CheckAuth(
		http.HandlerFunc(goodHandler.GetGoodByID),
	))
	mux.Handle("POST /goods", authMiddleware.CheckAuth(
		http.HandlerFunc(goodHandler.AddGood),
	))
	mux.Handle("PUT /goods/{good_id}", authMiddleware.CheckAuth(
		http.HandlerFunc(goodHandler.UpdateGood),
	))
	mux.Handle("DELETE /goods/{good_id}", authMiddleware.CheckAuth(
		http.HandlerFunc(goodHandler.DeleteGoodByID),
	))

	mux.HandleFunc("POST /auth/sign_up", authHandler.SignUp)
	mux.HandleFunc("POST /auth/sign_in", authHandler.SignIn)
	mux.HandleFunc("POST /auth/refresh", authHandler.RefreshTokens)

	mux.Handle("GET /docs/", httpSwagger.WrapHandler)

	return mux
}

func createGormDB() (*gorm.DB, error) {
	dsn := "host=postgres user=user password=user dbname=ecom_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
