package purchase

import (
	"context"
	"testing"
)

func TestPurchase(t *testing.T) {
	t.Run("scaffold", func(t *testing.T) {
		ctx := context.Background()

		if ctx == nil {
			t.Fatal("expected context")
		}
	})
}
