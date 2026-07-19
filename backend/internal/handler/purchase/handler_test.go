package purchase

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	purchasemock "github.com/JCKFinland/jck-connect/backend/internal/usecase/purchase/mock"
	"github.com/JCKFinland/jck-connect/backend/internal/middleware"
	purchaseusecase "github.com/JCKFinland/jck-connect/backend/internal/usecase/purchase"
)

func TestPurchase_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//--------------------------------------------------
	// Arrange
	//--------------------------------------------------

	userID := uuid.New()
	productID := uuid.New()

	mockService := new(purchasemock.Service)

	mockService.
		On(
			"Purchase",
			mock.Anything,
			purchaseusecase.PurchaseRequest{
				UserID:    userID,
				ProductID: productID,
			},
		).
		Return(nil)

	handler := New(mockService)

	body := []byte(`{
		"product_id":"` + productID.String() + `"
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/purchases",
		bytes.NewBuffer(body),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	c.Request = req

	c.Set(
		middleware.ContextUserID,
		userID.String(),
	)

	//--------------------------------------------------
	// Act
	//--------------------------------------------------

	handler.Purchase(c)

	//--------------------------------------------------
	// Assert
	//--------------------------------------------------

	assert.Equal(
		t,
		http.StatusCreated,
		w.Code,
	)

	mockService.AssertExpectations(t)
}