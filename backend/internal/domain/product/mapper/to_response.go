package mapper

import (
	productdto "github.com/JCKFinland/jck-connect/backend/internal/domain/product/dto"
	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
)

// ToResponse converts a Product entity into an API response.
func ToResponse(
	product *productentity.Product,
) *productdto.ProductResponse {

	if product == nil {
		return nil
	}

	return &productdto.ProductResponse{
		ID:        product.ID,
		Code:      product.Code,
		Name:      product.Name,
		Category:  product.Category,
		Provider:  product.Provider,
		Price:     product.Price,
		Currency:  product.Currency,
		Active:    product.Active,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

// ToResponseList converts Product entities into API responses.
func ToResponseList(
	products []*productentity.Product,
) []*productdto.ProductResponse {

	responses := make(
		[]*productdto.ProductResponse,
		0,
		len(products),
	)

	for _, product := range products {
		responses = append(
			responses,
			ToResponse(product),
		)
	}

	return responses
}
