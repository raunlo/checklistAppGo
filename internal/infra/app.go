package infra

import (
	"github.com/gin-gonic/gin"
)

type Application struct {
	v1Route HttpRouterV1
	router  *gin.Engine
	config  ServerConfiguration
}

func CreateApplication(v1 HttpRouterV1, router *gin.Engine, configuration ServerConfiguration) Application {
	return Application{
		v1Route: v1,
		router:  router,
		config:  configuration,
	}
}

func (application Application) StartApplication() error {
	routeApi := application.router.Group("/api")
	application.v1Route.CreateRoutes(routeApi)

	err := application.router.Run(application.config.port)
	return err
}
