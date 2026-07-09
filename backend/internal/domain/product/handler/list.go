package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	productmapper "github.com/JCKFinland/jck-connect/backend/internal/domain/product/mapper"
)

// List returns all products.
func (h *Handler) List(
	c *gin.Context,
) {

	products, err := h.service.List(
		c.Request.Context(),
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		productmapper.ToResponseList(products),
	)
}