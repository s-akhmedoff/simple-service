package rest

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service/models"
	"simple-service/utils"
	"simple-service/views"
	"time"
)

type productParams struct {
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
// @Param param body productParams true "product params"
// @Success 200 {object} views.R
// @Failure 400 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
// @Router /products [post]
func (s *Server) createProduct(c *gin.Context) {
	var req productParams

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
	productType := c.Param("type")
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
	ski := c.Param("ski")
	if ski == "" {
		s.log.Error("product ski is missing")
		s.reportError(c, http.StatusBadRequest, utils.RequestValidationErr, errors.New("product ski is missing"))

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

// readProductByID godoc
// @Summary Read Product By ID
// @Produce  json
// @Param id path string true "product's id"
// @Success 200 {object} views.R
// @Failure 400 {object} views.R
// @Failure 500 {object} views.R
// @Router /products/{id} [get]
func (s *Server) readProductByID(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		s.log.Error("product id is missing")
		s.reportError(c, http.StatusBadRequest, utils.RequestValidationErr, errors.New("product id is missing"))

		return
	}

	product, err := s.storage.Product().ReadByID(ID)
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusInternalServerError, utils.StorageOperationErr, err)

		return
	}

	s.respond(c, product)
}

// updateProduct godoc
// @Summary Update existing product
// @Accept  json
// @Produce  json
// @Param id path string true "product's ID"
// @Param param body productParams true "product params"
// @Success 200 {object} views.R
// @Failure 400 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
// @Router /products [put]
func (s *Server) updateProduct(c *gin.Context) {
	productID := c.Param("id")

	if productID == "" {
		s.log.Error("product ID is missing")
		s.reportError(c, http.StatusBadRequest, utils.RequestValidationErr, errors.New("product ID is missing"))

		return
	}

	var req productParams

	err := c.ShouldBindJSON(&req)
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusBadRequest, utils.RequestValidationErr, err)

		return
	}

	product, err := s.storage.Product().ReadByID(productID)
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusInternalServerError, utils.StorageOperationErr, err)

		return
	}

	now := time.Now()

	{
		product.ProductType = req.ProductType
		product.URI = req.URI
		product.Name = req.Name
		product.Description = req.Description
		product.IsActive = req.IsActive
		product.SKU = req.SKI
		product.UpdatedAt = &now
	}

	s.log.Debug(fmt.Sprintf("New product model to pass to storage: %+v", product))

	tx, err := s.storage.Tx()
	if err != nil {
		s.log.Error(err)
		s.reportError(c, http.StatusUnprocessableEntity, utils.StorageTxErr, err)

		return
	}

	defer func() { _ = tx.Rollback() }()

	err = s.storage.Product().Update(tx, productID, product)
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

	s.respond(c, product)
}

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
