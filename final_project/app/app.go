package app

import (
	"log"
	"net/http"
	"time"

	"ecom/internal/apicollector"
	"ecom/internal/converter"
	"ecom/internal/domain"
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

	apiCollector := apicollector.NewApiCollector()
	apiCollectorMiddleware := middleware.NewApiCollectorMiddleware(apiCollector)

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

	mux := newServeMux(goodHandler, authHandler, authMiddleware, apiCollectorMiddleware)

	srv := http.Server{
		Addr:    serverPort,
		Handler: mux,
	}

	doneCh := make(chan struct{})
	go runApiInfoPrint(time.Now(), apiCollector, doneCh)

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func newServeMux(
	goodHandler *rest.GoodHandler,
	authHandler *rest.AuthHandler,
	authMiddleware *middleware.AuthMiddleware,
	apiCollectorMiddleware *middleware.ApiCollectorMiddleware,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /goods", authMiddleware.CheckAuth(
		apiCollectorMiddleware.CollectInfo(
			http.HandlerFunc(goodHandler.GetAllGoods),
		),
	))
	mux.Handle("GET /goods/{good_id}", authMiddleware.CheckAuth(
		apiCollectorMiddleware.CollectInfo(
			http.HandlerFunc(goodHandler.GetGoodByID),
		),
	))
	mux.Handle("POST /goods", authMiddleware.CheckRole(
		apiCollectorMiddleware.CollectInfo(
			http.HandlerFunc(goodHandler.AddGood),
		),
		domain.AdminRole,
	))
	mux.Handle("PUT /goods/{good_id}", authMiddleware.CheckRole(
		apiCollectorMiddleware.CollectInfo(
			http.HandlerFunc(goodHandler.UpdateGood),
		),
		domain.AdminRole,
	))
	mux.Handle("DELETE /goods/{good_id}", authMiddleware.CheckRole(
		apiCollectorMiddleware.CollectInfo(
			http.HandlerFunc(goodHandler.DeleteGoodByID),
		),
		domain.AdminRole,
	))

	mux.Handle("POST /auth/sign_up", apiCollectorMiddleware.CollectInfo(
		http.HandlerFunc(authHandler.SignUp),
	))
	mux.Handle("POST /auth/sign_in", apiCollectorMiddleware.CollectInfo(
		http.HandlerFunc(authHandler.SignIn),
	))
	mux.Handle("POST /auth/refresh", apiCollectorMiddleware.CollectInfo(
		http.HandlerFunc(authHandler.RefreshTokens),
	))

	mux.Handle("GET /docs/", apiCollectorMiddleware.CollectInfo(httpSwagger.WrapHandler))

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

func runApiInfoPrint(startTime time.Time,
	apiCollector apicollector.ApiCollector,
	done <-chan struct{},
) {
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			log.Printf("Time from start service: %v", time.Since(startTime))
			apiCollector.PrintApiInfo()
		}
	}
}
