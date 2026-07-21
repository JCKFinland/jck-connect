package mapper

import (
	productdto "github.com/JCKFinland/jck-connect/backend/internal/domain/product/dto"
	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
)

// CreateRequestToEntity converts a create request into a Product entity.
func CreateRequestToEntity(
	request *productdto.CreateProductRequest,
) *productentity.Product {

	return &productentity.Product{
		Code:     request.Code,
		Name:     request.Name,
		Category: request.Category,
		Provider: request.Provider,
		Price:    request.Price,
		Currency: request.Currency,
	}
}

// UpdateRequestToEntity converts an update request into a Product entity.
func UpdateRequestToEntity(
	request *productdto.UpdateProductRequest,
) *productentity.Product {

	return &productentity.Product{
		Code:     request.Code,
		Name:     request.Name,
		Category: request.Category,
		Provider: request.Provider,
		Price:    request.Price,
		Currency: request.Currency,
		Active:   request.Active,
	}
}
