package router

import (
	"github.com/figarocms/hr-go-utils/v2/logger"
	fcmsrouter "github.com/figarocms/hr-go-utils/v2/logger/gin"
	"github.com/gin-gonic/gin"
	"github.com/vdamery/jdria/internal/api/handlers"
	"github.com/vdamery/jdria/internal/api/services"
	"time"
)

func InitializeRouter(services *services.Services) *gin.Engine {
	r := gin.New()

	// Initialize Global Middlewares
	initializeMiddlewares(r)

	// Initialize routes
	initializeRoutes(r, services)

	return r
}

func initializeRoutes(r *gin.Engine, services *services.Services) {
	r.GET("/health", handlers.Health())

	r.GET("/", handlers.Index(services.GamerService))
	r.POST("/send", handlers.Send(services.GamerService))

	r.Use(fcmsrouter.Ginzap(logger.Log, time.RFC3339, true, false))
	r.NoRoute(handlers.NoRoute)
	r.LoadHTMLGlob("**/templates/*")
}
