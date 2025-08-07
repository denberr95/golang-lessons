package api

import (
	"context"
	"main/config"
	"main/logging"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var router *gin.Engine
var log *logrus.Logger

func Init(cfg *config.WebServerConfig) {
	router = gin.Default()
	log = logging.GetLogger()

	addRoutes(cfg)

	s := &http.Server{
		Addr:           cfg.GetFullAddress(),
		Handler:        router,
		ReadTimeout:    time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.WriteTimeout) * time.Second,
		MaxHeaderBytes: cfg.MaxHeaderSizeMB * 1024 * 1024,
		IdleTimeout:    time.Duration(cfg.IdleTimeout) * time.Second,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting Web Server: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Web Server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.GracefulShutdownTime)*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Println("Error Shutdown Web Server: ", err)
	}
	log.Println("Server stopped !")
}

func addRoutes(cfg *config.WebServerConfig) {
	registerUserRoutes(cfg.BasePath, router)
}

func GetRouter() *gin.Engine {
	return router
}
