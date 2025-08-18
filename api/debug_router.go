package api

import "github.com/gin-gonic/gin"

func configureDebugRouterLogFunc() {
	if webServerConfig.Log.EnablePrintExposedRouter {
		gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
			log.Infof("Servizio esposto su httpMethod: %v - path: %v - handler: %v - nuHandlers: %v", httpMethod, absolutePath, handlerName, nuHandlers)
		}
	}
}
