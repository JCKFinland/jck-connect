package service

import (
	"fmt"
	"time"
)

// generateReference generates a unique transaction reference.
func generateReference() string {
	return fmt.Sprintf(
		"TXN-%d",
		time.Now().UTC().UnixNano(),
	)
}