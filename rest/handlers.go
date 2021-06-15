package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"simple-service/views"
)

func (s *Server) handleSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, views.R{
		Status:    "Success",
		ErrorCode: 0,
		ErrorNote: "",
		Data:      data,
	})
}

func (s *Server) handleErrorResponse(c *gin.Context, httpCode, errorCode int, err error) {
	c.JSON(httpCode, views.R{
		Status:    "Failure",
		ErrorCode: errorCode,
		ErrorNote: err.Error(),
		Data:      nil,
	})
}

// config godoc
// @Summary Show config
// @Description expose all configs for debugging
// @Produce json
// @Success 200 {object} views.R{data=config.Config}
// @Failure 405 {object} views.R
// @Header 200,405 {string} environment "Current environment"
// @Header 200,405 {string} go-version "Version of Golang"
// @Header 200,405 {string} go-os "Go OS"
// @Router /config [get]
func (s *Server) config(c *gin.Context) {
	s.log.Debug("handle /config endpoint")

	c.Header("environment", s.cfg.Environment)
	c.Header("go-version", runtime.Version())
	c.Header("go-os", runtime.GOOS)

	if s.cfg.Environment == "dev" {
		s.handleSuccessResponse(c, s.cfg)

		return
	}

	s.handleErrorResponse(c, http.StatusMethodNotAllowed, -100, errors.New("deployed staging/production environment"))
}
