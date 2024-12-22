package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"wallet/cmd/api/controller"
	"wallet/cmd/api/service"
	"wallet/cmd/api/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	port   string
	DBPass string
	DBName string
	DBUser string
}

func main() {
	cfg := initConfig()

	dns := fmt.Sprintf(
		"host=localhost user=%s password=%s  dbname=%s port=5432 sslmode=disable",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var (
		str    = storage.NewStorage(db)
		svc    = service.NewService(str)
		ctrl   = controller.NewAPIController(svc)
		router = gin.Default()
	)

	router.Use(controller.ErrorHandler())

	ctrl.SetupRouter(router)

	srv := http.Server{
		Addr:    ":" + cfg.port,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func initConfig() *config {
	var cfg config
	var ok bool

	if cfg.port = os.Getenv("PORT"); cfg.port == "" {
		cfg.port = "8080"
	}

	if cfg.DBName, ok = os.LookupEnv("POSTGRES_DB"); !ok {
		log.Fatal("POSTGRES_DB environment variable must be set")
	}

	if cfg.DBUser, ok = os.LookupEnv("POSTGRES_USER"); !ok {
		log.Fatal("POSTGRES_USER environment variable must be set")
	}

	if cfg.DBPass, ok = os.LookupEnv("POSTGRES_PASSWORD"); !ok {
		log.Fatal("POSTGRES_PASSWORD environment variable must be set")
	}

	return &cfg
}
