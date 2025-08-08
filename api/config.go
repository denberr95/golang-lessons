package api

import (
	"context"
	"fmt"
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
var log *logrus.Logger = logging.GetLogger()
var cfg *config.WebServerConfig = &config.GetConfig().GoApp.WebServer

func GetRouter() *gin.Engine {
	return router
}

func Run() {
	router = gin.New()
	router.MaxMultipartMemory = cfg.GetMaxHeaderSizeMB()
	router.Use(gin.Recovery())

	configureLogMiddleware()
	configureAccessLogs()
	configureDebugRouterLogFunc()
	addRoutes()
	runWebServer()
}

func configureLogMiddleware() {
	if cfg.EnableLogMiddleware {
		router.Use(logMiddleware())
	}
}

func configureAccessLogs() {
	if cfg.EnableAccessLog {
		router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("[%s] - [%s] - [%s] - [%s] - [%s] - [%d] - [%s] [%s] - [%s]",
				param.TimeStamp.Format(time.RFC3339),
				param.ClientIP,
				param.Method,
				param.Request.RequestURI,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		}))
	}
}

func configureDebugRouterLogFunc() {
	if cfg.EnablePrintExposedRouter {
		gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
			log.Infof("Servizio esposto su httpMethod=%v, path=%v handler=%v nuHandlers=%v", httpMethod, absolutePath, handlerName, nuHandlers)
		}
	}
}

func addRoutes() {
	router.Group(cfg.BasePath)
	registerUserRoutes(router)
}

func runWebServer() {
	s := &http.Server{
		Handler:        router,
		Addr:           cfg.GetFullAddress(),
		ReadTimeout:    cfg.GetReadTimeout(),
		WriteTimeout:   cfg.GetWriteTimeout(),
		MaxHeaderBytes: cfg.GetMaxHeaderBytes(),
		IdleTimeout:    cfg.GetIdleTimeout(),
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Errore avvio Web Server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Stop Web Server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.GracefulShutdownTime)*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Println("Errore Stop Web Server: ", err)
	}
	log.Println("Web Server stoppato !")
}
