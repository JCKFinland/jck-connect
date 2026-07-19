package purchase

import (

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	purchaserequest "github.com/JCKFinland/jck-connect/backend/internal/request/purchase"
	purchaseusecase "github.com/JCKFinland/jck-connect/backend/internal/usecase/purchase"
	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"

	"github.com/JCKFinland/jck-connect/backend/internal/middleware"
)

// Purchase handles purchase requests.
//
// Implementation will be completed in the next batches.
func (h *Handler) Purchase(
	c *gin.Context,
) {

	//--------------------------------------------------
	// Parse request
	//--------------------------------------------------

	var request purchaserequest.PurchaseRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(
			sharedErrors.HTTPStatus(err),
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	//--------------------------------------------------
	// Authenticated user
	//--------------------------------------------------

	userID, err := uuid.Parse(
		middleware.UserID(c),
	)
	if err != nil {

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "invalid authenticated user",
			},
		)

		return
	}

	//--------------------------------------------------
// Execute purchase
//--------------------------------------------------

err = h.purchaseService.Purchase(
    c.Request.Context(),
    purchaseusecase.PurchaseRequest{
        UserID:    userID,
        ProductID: request.ProductID,
    },
)

if err != nil {

    c.JSON(
        sharedErrors.HTTPStatus(err),
        gin.H{
            "code":    sharedErrors.Code(err),
            "message": sharedErrors.Message(err),
        },
    )

    return
}

	//--------------------------------------------------
	// Success
	//--------------------------------------------------

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "purchase completed successfully",
		},
	)
}