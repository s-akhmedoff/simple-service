package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		Status:    "Success",
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
// @Header 200 {string} Environment "dev"
// @Router /config [get]
func (s *Server) config(c *gin.Context) {
	s.log.Debug("handle /config endpoint")
	s.handleSuccessResponse(c, s.cfg)
}
