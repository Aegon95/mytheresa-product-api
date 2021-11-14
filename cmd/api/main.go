package main

import (
	"context"
	"fmt"
	"github.com/Aegon95/mytheresa-product-api/config"
	"github.com/Aegon95/mytheresa-product-api/internal/db"
	"github.com/Aegon95/mytheresa-product-api/internal/server"
	constants "github.com/Aegon95/mytheresa-product-api/internal/util"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	zapLogger, err := zap.NewProduction()
	logger := zapLogger.Sugar()
	if err != nil {
		log.Fatalf("zap.NewProduction %s", err)
	}

	logger.Debug("Reading the Config file")

	cfg := config.Setup(logger, constants.CONFIG_PATH)

	logger.Info("Finished reading Config file")

	logger.Debug("Opening database connection")

	database := db.NewDatabase(logger, cfg)

	pg := database.Setup()

	logger.Info("Successfully opened database connection")

	logger.Debug("Setting up Routers")

	router := server.Setup(logger, pg)

	logger.Info("Finished setting up Routers")

	srv := http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	startServer(&srv)
}

func startServer(server *http.Server) {

	serverCtx, serverStopCtx := context.WithCancel(context.Background())


	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig


		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()


		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()


	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}


	<-serverCtx.Done()
}

