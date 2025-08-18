package api

import (
	"context"
	"main/config"
	"main/logging"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var router *gin.Engine
var log *logrus.Logger = logging.GetLogger()
var webServerConfig *config.WebServerConfig = &config.GetConfig().GoApp.WebServer

func GetRouter() *gin.Engine {
	return router
}

func Run() {
	router = gin.New()
	router.MaxMultipartMemory = webServerConfig.HTTP.GetMaxHeaderSizeMB()
	router.Use(gin.Recovery())

	configureLogMiddleware()
	configureAccessLog()
	configureDebugRouterLogFunc()
	addRoutes()
	runWebServer()
}

func configureLogMiddleware() {
	if webServerConfig.Log.EnableLogMiddleware {
		router.Use(logMiddleware())
	}
}

func addRoutes() {
	var baseGroup = router.Group(webServerConfig.Base.BasePath)
	registerUserRoutes(baseGroup)
}

func runWebServer() {
	s := &http.Server{
		Handler:        router,
		Addr:           webServerConfig.Base.GetFullAddress(),
		ReadTimeout:    webServerConfig.HTTP.ReadTimeout,
		WriteTimeout:   webServerConfig.HTTP.WriteTimeout,
		MaxHeaderBytes: webServerConfig.HTTP.GetMaxHeaderBytes(),
		IdleTimeout:    webServerConfig.HTTP.IdleTimeout,
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

	ctx, cancel := context.WithTimeout(context.Background(), webServerConfig.HTTP.GracefulShutdownTime)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Println("Errore Stop Web Server: ", err)
	}
	log.Println("Web Server stoppato !")
}
