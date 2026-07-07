package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success sends a standard HTTP 200 response.
func Success(
	c *gin.Context,
	message string,
	data interface{},
) {
	c.JSON(
		http.StatusOK,
		Response{
			Success: true,
			Message: message,
			Data:    data,
		},
	)
}

// Created sends a standard HTTP 201 response.
func Created(
	c *gin.Context,
	message string,
	data interface{},
) {
	c.JSON(
		http.StatusCreated,
		Response{
			Success: true,
			Message: message,
			Data:    data,
		},
	)
}

// NoContent sends a standard HTTP 204 response.
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}