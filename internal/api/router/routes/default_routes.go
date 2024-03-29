package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/oprimogus/cardapiogo/docs"
	"github.com/oprimogus/cardapiogo/internal/domain/factory"
)

func DefaultRoutes(router *gin.Engine, factory factory.RepositoryFactory, reg *prometheus.Registry) {

	basePath := os.Getenv("API_BASE_PATH")

	v1 := router.Group(basePath + "/v1")

	// SWAGGER
	docs.SwaggerInfo.BasePath = basePath
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// Prometheus
	v1.GET("/metrics", func(c *gin.Context) {
		h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
		h.ServeHTTP(c.Writer, c.Request)
	})

}
