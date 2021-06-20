package rest

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) endpoints() {
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.router.GET("/config", s.config)

	s.router.POST("/products", s.createProduct)
	s.router.GET("/products", s.readProducts)
	s.router.GET("/products/:type", s.readProductByType)
	s.router.GET("/products/:ski", s.readProductBySKI)
	s.router.GET("/products/:id", s.readProductByID)
	s.router.PUT("/products/:id", s.updateProduct)
	s.router.DELETE("/products/:id", s.deleteProduct)
}
