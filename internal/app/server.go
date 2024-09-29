package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go-react-app/internal/app/api"
	"go-react-app/internal/app/api/health"
	"go-react-app/internal/app/api/hello"
	"go-react-app/internal/app/generated"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Server struct {
	SPAPath       string
	ClientOrigins []string
	Health        *health.Controller
	Hello         *hello.Controller
}

func NewServer(spaPath string, clientOrigins []string) *Server {
	return &Server{
		SPAPath:       spaPath,
		ClientOrigins: clientOrigins,
		Health:        health.NewHealth(),
		Hello:         hello.NewHello(),
	}
}

func (s *Server) StartServer() error {
	router := gin.New()

	// Since we don't use any proxy, this feature can be disabled
	err := router.SetTrustedProxies(nil)
	if err != nil {
		return fmt.Errorf("could not set trusted proxies: %v", err)
	}

	router.Use(newGinLogger())
	router.Use(newCorsHandler(s.ClientOrigins))
	router.Use(errorHandler)
	newPrometheus().Use(router)

	// API ROUTES
	generated.RegisterHandlersWithOptions(
		router,
		api.NewServer(),
		generated.GinServerOptions{
			BaseURL:      "/api",
			ErrorHandler: nil,
		})

	// SPA ROUTE
	// Only loaded if SPAPath is defined.
	if s.SPAPath != "" {
		router.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusPermanentRedirect, "/app")
		})

		log.Debug().Str("spaPath", s.SPAPath).Msg("SPA_PATH is set, will serve")

		router.Static("/app", s.SPAPath)
	}

	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/app") && s.SPAPath != "" {
			c.File(fmt.Sprintf("%s/index.html", s.SPAPath))
			return
		}

		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	return router.Run()
}

func newCorsHandler(clientOrigins []string) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = clientOrigins
	if len(corsConfig.AllowOrigins) > 0 {
		log.Info().Interface("allowedOrigins", corsConfig.AllowOrigins).Msg("CORS origins configured")
	}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(
		corsConfig.AllowHeaders,
		[]string{
			"Authorization",
		}...,
	)

	return cors.New(corsConfig)
}

func newPrometheus() *ginprometheus.Prometheus {
	prometheus := ginprometheus.NewPrometheus("gin")

	// Prevents high cardinality of metrics Source: https://github.com/zsais/go-gin-prometheus#preserving-a-low-cardinality-for-the-request-counter
	prometheus.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL.Path // Query params are dropped here so there is not a metric for every permutation of query param usage on a route

		//  If a route uses parameters, replace the parameter value with its name. Else there will be a metric for the route
		//  with every possible value of that parameter and this will cause performance issues in Prometheus.
		//
		//  If your service uses route parameters, uncomment the for loop below and add a case for each parameter. The example case
		//  below works for routes with a parameter called 'name', like '/api/function/:name'
		//  --
		//    for _, p := range c.Params {
		//      switch p.Key {
		//      case "name":
		//        url = strings.Replace(url, p.Value, ":name", 1)
		//      }
		//    }
		return url
	}

	return prometheus
}

func newGinLogger() gin.HandlerFunc {
	return logger.SetLogger(
		logger.WithSkipPath([]string{
			"/health",
			"/metrics",
		}),
	)
}
