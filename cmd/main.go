package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"simple-service/api/docs"
	"simple-service/config"
	"simple-service/rest"
	"simple-service/storage/postgres"

	_ "github.com/lib/pq"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// *** Bootstrap
	cfg := config.Load()
	log := cfg.LoggerSetup()

	db, err := sqlx.Connect("postgres", cfg.PostgresURL())
	if err != nil {
		log.WithFields(logrus.Fields{"scope": "entrypoint:main", "env": cfg.Environment}).Fatal(err)
	}

	storage := postgres.NewStore(cfg, log, db)

	// *** Swagger Docs
	docs.SwaggerInfo.Host = cfg.HTTPHost + cfg.HTTPPort
	docs.SwaggerInfo.Title = config.ServiceName
	docs.SwaggerInfo.Description = "Simple service"

	// *** REST Initialization
	r := gin.Default()

	handler := rest.NewAPIServer(cfg, log, storage, r)

	srv := &http.Server{
		Addr:    cfg.HTTPPort,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithFields(logrus.Fields{"scope": "entrypoint:main", "env": cfg.Environment}).Fatal(fmt.Sprintf("Failed To Start REST Server: %s\n", err.Error()))
		}
	}()

	log.WithFields(logrus.Fields{"scope": "entrypoint:main", "env": cfg.Environment}).Info("REST API server started at " + cfg.HTTPHost + cfg.HTTPPort)

	<-quit
	log.Info("Interrupted")

	ctx, cancel := context.WithCancel(context.Background())

	if err := srv.Shutdown(ctx); err != nil {
		log.WithFields(logrus.Fields{"scope": "entrypoint:main", "env": cfg.Environment}).Error(fmt.Sprintf("REST Server Graceful Shutdown Failed: %s\n", err))
	}

	cancel()
}
