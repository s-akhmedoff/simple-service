package rest

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service/models"
	"simple-service/utils"
	"simple-service/views"
)

type createProductParams struct {
	SKI         string `json:"ski" binding:"required"`
	Name        string `json:"name" binding:"required"`
	ProductType string `json:"product_type" binding:"required"`
	URI         string `json:"uri" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsActive    bool   `json:"is_active" binding:"required"`
}

// createProduct godoc
// @Summary Create new product
// @Description create new product in storage
// @Accept  json
// @Produce  json
// @Param param body createProductParams true "product params"
// @Success 200 {object} views.R
// @Failure 400 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
// @Router /products [post]
func (s *Server) createProduct(c *gin.Context) {
	var req createProductParams

	if err := c.ShouldBindJSON(&req); err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusBadRequest, utils.RequestValidationErr, err)

		return
	}

	s.log.Debug(fmt.Sprintf("Request successful binded: %+v", req))

	tx, err := s.storage.Tx()
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusUnprocessableEntity, utils.StorageTxErr, err)

		return
	}

	defer func() { _ = tx.Rollback() }()

	product := models.NewProduct{
		SKI:         req.SKI,
		Name:        req.Name,
		ProductType: req.ProductType,
		URI:         req.URI,
		Description: req.Description,
		IsActive:    req.IsActive,
	}

	s.log.Debug(fmt.Sprintf("New product model to pass to storage: %+v", product))

	err = s.storage.Product().Create(tx, &product)
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusInternalServerError, utils.StorageOperationErr, err)

		return
	}

	err = tx.Commit()
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusInternalServerError, utils.StorageTxErr, err)

		return
	}

	s.log.Debug("Product created, transaction was successfully committed")

	products, err := s.storage.Product().Read()
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusInternalServerError, utils.StorageOperationErr, err)

		return
	}

	c.JSON(http.StatusCreated, views.R{
		Status:    "Success",
		ErrorCode: 0,
		ErrorNote: "",
		Data: struct {
			Total int             `json:"total"`
			Items *models.Product `json:"items"`
		}{
			Total: 1,
			Items: products[len(products)-1],
		},
	})
}

// readProducts godoc
// @Summary Get all products
// @Produce  json
// @Success 200 {object} views.R
// @Failure 400 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
// @Router /products [get]
func (s *Server) readProducts(c *gin.Context) {
	products, err := s.storage.Product().Read()
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusInternalServerError, utils.StorageOperationErr, err)

		return
	}

	c.JSON(http.StatusCreated, views.R{
		Status:    "Success",
		ErrorCode: 0,
		ErrorNote: "",
		Data: struct {
			Total int               `json:"total"`
			Items []*models.Product `json:"items"`
		}{
			Total: len(products),
			Items: products,
		},
	})
}

// readProductByType godoc
// @Summary Read product by type
// @Description get product by its type
// @Produce  json
// @Param type path string true "product's type"
// @Success 200 {object} views.R
// @Failure 400 {object} views.R
// @Failure 500 {object} views.R
// @Router /products/{type} [post]
func (s *Server) readProductByType(c *gin.Context) {
	productType := c.Query("type")
	if productType == "" {
		s.log.Error("product type is missing")
		s.reportError(c, http.StatusBadRequest, utils.RequestValidationErr, errors.New("product type is missing"))

		return
	}

	product, err := s.storage.Product().ReadByType(productType)
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusInternalServerError, utils.StorageOperationErr, err)

		return
	}

	s.respond(c, product)
}

// readProductBySKI godoc
// @Summary Read Product By SKI
// @Description get product by its ski
// @Produce  json
// @Param ski path string true "product's ski"
// @Success 200 {object} views.R
// @Failure 400 {object} views.R
// @Failure 500 {object} views.R
// @Router /products/{ski} [get]
func (s *Server) readProductBySKI(c *gin.Context) {
	ski := c.Query("type")
	if ski == "" {
		s.log.Error("product type is missing")
		s.reportError(c, http.StatusBadRequest, utils.RequestValidationErr, errors.New("product type is missing"))

		return
	}

	product, err := s.storage.Product().ReadBySKI(ski)
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusInternalServerError, utils.StorageOperationErr, err)

		return
	}

	s.respond(c, product)
}

// updateProduct godoc
// @Summary Create new product
// @Description create new product in storage
// @Accept  json
// @Produce  json
// @Param param body createProductParams true "product params"
// @Success 200 {object} views.R
// @Failure 400 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
// @Router /products [put]
//func (s *Server) updateProduct(c *gin.Context) {}

// deleteProduct godoc
// @Summary Delete product
// @Produce  json
// @Param id path string true "product's ID"
// @Success 200 {object} views.R
// @Failure 400 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
// @Router /products/{id} [delete]
func (s *Server) deleteProduct(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		s.log.Error("product ID is missing")
		s.reportError(c, http.StatusBadRequest, utils.RequestValidationErr, errors.New("product ID is missing"))

		return
	}

	tx, err := s.storage.Tx()
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusUnprocessableEntity, utils.StorageTxErr, err)

		return
	}

	defer func() { _ = tx.Rollback() }()

	err = s.storage.Product().Delete(tx, id)
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusInternalServerError, utils.StorageOperationErr, err)

		return
	}

	err = tx.Commit()
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusInternalServerError, utils.StorageTxErr, err)

		return
	}

	s.respond(c, nil)
}
