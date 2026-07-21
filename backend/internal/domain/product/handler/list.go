package handler

import (
	"github.com/gin-gonic/gin"

	productmapper "github.com/JCKFinland/jck-connect/backend/internal/domain/product/mapper"
	response "github.com/JCKFinland/jck-connect/backend/internal/shared/response"
)

// List returns all products.
func (h *Handler) List(
	c *gin.Context,
) {

	products, err := h.service.List(
		c.Request.Context(),
	)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.Success(
		c,
		response.MsgProductsRetrieved,
		productmapper.ToResponseList(products),
	)
}
