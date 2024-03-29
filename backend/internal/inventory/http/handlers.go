package http

import (
	"firdausyusofs/kem_digital/config"
	"firdausyusofs/kem_digital/internal/inventory"
	"firdausyusofs/kem_digital/internal/models"
	"firdausyusofs/kem_digital/pkg/api_response"
	"firdausyusofs/kem_digital/pkg/logger"
	"firdausyusofs/kem_digital/pkg/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type inventoryHandler struct {
	cfg         *config.Config
	inventoryUC inventory.UseCase
	logger      logger.Logger
}

func NewInventoryHandler(cfg *config.Config, inventoryUC inventory.UseCase, logger logger.Logger) inventory.Handlers {
	return &inventoryHandler{
		cfg:         cfg,
		inventoryUC: inventoryUC,
		logger:      logger,
	}
}

func (h *inventoryHandler) GetProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Initiate pagination query
		pq, err := utils.NewPaginationQuery(c)
		if err != nil {
			errMsg := err.Error()
			return c.JSON(api_response.MakeErrorResponse(http.StatusBadRequest, "Invalid pagination query", &errMsg))
		}

		products, err := h.inventoryUC.GetProducts(c.Request().Context(), pq)
		if err != nil {
			errMsg := err.Error()
			return c.JSON(api_response.MakeErrorResponse(http.StatusNotFound, "Failed to retrieve products", &errMsg))
		}

		return c.JSON(api_response.MakeSuccessResponse(http.StatusOK, "Successfully retrieved products", &products))
	}
}

func (h *inventoryHandler) GetProductByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		productId, err := strconv.ParseInt(c.Param("product_id"), 10, 64)
		if err != nil {
			errMsg := err.Error()
			return c.JSON(api_response.MakeErrorResponse(http.StatusBadRequest, "Invalid product ID", &errMsg))
		}

		product, err := h.inventoryUC.GetProductByID(c.Request().Context(), productId)
		if err != nil {
			errMsg := err.Error()
			return c.JSON(api_response.MakeErrorResponse(http.StatusNotFound, "Failed to retrieve product", &errMsg))
		}

		return c.JSON(api_response.MakeSuccessResponse(http.StatusOK, "Successfully retrieved product", &product))
	}
}

func (h *inventoryHandler) CreateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		product := &models.Product{}
		if err := utils.ReadBody(c, product); err != nil {
			errMsg := err.Error()
			return c.JSON(api_response.MakeErrorResponse(http.StatusBadRequest, "Invalid request body", &errMsg))
		}

		createdProduct, err := h.inventoryUC.CreateProduct(c.Request().Context(), product)
		if err != nil {
			errMsg := err.Error()
			return c.JSON(api_response.MakeErrorResponse(http.StatusInternalServerError, "Failed to create product", &errMsg))
		}

		return c.JSON(api_response.MakeSuccessResponse(http.StatusCreated, "Successfully created product", &createdProduct))
	}
}

func (h *inventoryHandler) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		productID, err := strconv.ParseInt(c.Param("product_id"), 10, 64)
		if err != nil {
			return c.JSON(api_response.MakeErrorResponse(http.StatusBadRequest, "Invalid product ID", nil))
		}

		product := &models.Product{}
		product.ID = uint(productID)

		if err := utils.ReadBody(c, product); err != nil {
			return c.JSON(api_response.MakeErrorResponse(http.StatusBadRequest, "Invalid request body", nil))
		}

		updatedProduct, err := h.inventoryUC.UpdateProduct(c.Request().Context(), product)
		if err != nil {
			return c.JSON(api_response.MakeErrorResponse(http.StatusInternalServerError, "Failed to update product", nil))
		}

		return c.JSON(api_response.MakeSuccessResponse(http.StatusOK, "Successfully updated product", &updatedProduct))
	}
}

func (h *inventoryHandler) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		productID, err := strconv.ParseInt(c.Param("product_id"), 10, 64)
		if err != nil {
			return c.JSON(api_response.MakeErrorResponse(http.StatusBadRequest, "Invalid product ID", nil))
		}

		err = h.inventoryUC.DeleteProduct(c.Request().Context(), productID)
		if err != nil {
			return c.JSON(api_response.MakeErrorResponse(http.StatusInternalServerError, "Failed to delete product", nil))
		}

		return c.JSON(api_response.MakeSuccessResponse(http.StatusOK, "Successfully deleted product", nil))
	}
}
