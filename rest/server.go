package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zsais/go-gin-prometheus"
	"net/http"
	"simple-service/config"
	"simple-service/storage"
	"strings"
)

const scope = "server:rest_api"

// Server - API Server
// @title Simple Service REST API
// @version 1.0
// @description This is a sample server
// @contact.name s-akhmedoff
// @contact.url https://github.com/s-akhmedoff
// @contact.email sodikjon.akhmedoff@gmail.com
// @license.name GNU GPLv3
// @license.url https://www.gnu.org/licenses/gpl-3.0.en.html
// @host localhost:7070
type Server struct {
	cfg     *config.Config
	log     *logrus.Entry
	storage storage.Storage
	router  *gin.Engine
}

// NewAPIServer - wrapper on gin engine, sets up endpoints
func NewAPIServer(cfg *config.Config, log *logrus.Logger, storage storage.Storage, router *gin.Engine) (s *Server) {
	s = &Server{
		cfg:     cfg,
		log:     log.WithFields(logrus.Fields{"scope": scope, "env": cfg.Environment}),
		storage: storage,
		router:  router,
	}

	s.metrics() // Bind prometheus metrics exporter

	s.endpoints() // Binds all endpoints

	return
}

// ServerHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) metrics() {
	p := ginprometheus.NewPrometheus("gin")

	p.Use(s.router)

	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL.Path
		for _, p := range c.Params {
			if p.Key == "name" {
				url = strings.Replace(url, p.Value, ":name", 1)
				break
			}
		}
		return url
	}

}
