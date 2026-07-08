package mapper

import (
	orderdto "github.com/JCKFinland/jck-connect/backend/internal/domain/order/dto"
	orderentity "github.com/JCKFinland/jck-connect/backend/internal/domain/order/entity"
)

// ToOrderResponse converts an Order entity into an API response DTO.
func ToOrderResponse(
	order *orderentity.Order,
) *orderdto.OrderResponse {

	if order == nil {
		return nil
	}

	return &orderdto.OrderResponse{
		ID:        order.ID,
		UserID:    order.UserID,
		ProductID: order.ProductID,
		Reference: order.Reference,
		Amount:    order.Amount,
		Currency:  order.Currency,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}

// ToOrderResponseList converts a slice of Order entities into response DTOs.
func ToOrderResponseList(
	orders []*orderentity.Order,
) []*orderdto.OrderResponse {

	if len(orders) == 0 {
		return []*orderdto.OrderResponse{}
	}

	responses := make([]*orderdto.OrderResponse, 0, len(orders))

	for _, order := range orders {
		responses = append(
			responses,
			ToOrderResponse(order),
		)
	}

	return responses
}
