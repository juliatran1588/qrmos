package controller

import (
	"log"
	"net/http"
	"qrmos/internal/adapter/controller/internal/authcheck"
	"qrmos/internal/common/config"
	"qrmos/internal/usecase/repo"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	UserRepo     repo.User
	DeliveryRepo repo.Delivery
}

func NewServer(cfg ServerConfig) *server {
	if config.App().ENV != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	return &server{
		r:            gin.Default(),
		userRepo:     cfg.UserRepo,
		deliveryRepo: cfg.DeliveryRepo,
		authCheck:    authcheck.NewAuthCheck(cfg.UserRepo),
	}
}

type server struct {
	r *gin.Engine

	userRepo     repo.User
	deliveryRepo repo.Delivery

	authCheck *authcheck.AuthCheck
}

func (s *server) Run() {
	switch config.App().ENV {
	case "dev":
		s.runDev()
	case "staging":
		s.runStaging()
	case "prod":
		s.runProd()
	default:
		panic("unsupported environment")
	}
}

func (s *server) runDev() {
	s.r.Use(cors.Default())
	s.setupAPIs()

	port := strconv.Itoa(config.App().Port)
	s.r.Run(":" + port)
}

func (s *server) runStaging() {
	s.serveWebStatic()
	s.setupAPIs()

	port := strconv.Itoa(config.App().Port)
	s.r.Run(":" + port)
}

func (s *server) runProd() {
	s.serveWebStatic()
	s.setupAPIs()

	s.r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	log.Fatal(autotls.Run(s.r, config.App().Domains...))
}

func (s *server) serveWebStatic() {
	s.r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/web/")
	})
	s.serveStatic("/web/", "./web")
}

func (s *server) serveStatic(relativePath, root string) {
	s.r.Use(addCacheControlHeaderFor(relativePath))
	s.r.Static(relativePath, root)
}

func addCacheControlHeaderFor(prefix string) func(*gin.Context) {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, prefix) {
			c.Header("Cache-Control", "max-age=3600")
			c.Next()
		}
	}
}
