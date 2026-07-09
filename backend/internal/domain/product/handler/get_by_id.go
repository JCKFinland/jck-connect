package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	productmapper "github.com/JCKFinland/jck-connect/backend/internal/domain/product/mapper"
)

// GetByID returns a product by ID.
func (h *Handler) GetByID(
	c *gin.Context,
) {

	id, err := uuid.Parse(
		c.Param("id"),
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid product id",
			},
		)
		return
	}

	product, err := h.service.GetByID(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		productmapper.ToResponse(product),
	)
}