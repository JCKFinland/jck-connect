package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// generateReference creates a unique public order reference.
//
// Format:
//
//	JCK-YYYYMMDD-XXXXXXXX
//
// Example:
//
//	JCK-20260707-84F3A9D2
//
// This reference is intended for customers and external systems.
// It is independent of the internal UUID used as the primary key.
func generateReference() string {

	// Generate 4 random bytes (8 hexadecimal characters).
	random := make([]byte, 4)

	if _, err := rand.Read(random); err != nil {
		// Extremely unlikely. Fall back to a timestamp-derived suffix.
		return fmt.Sprintf(
			"JCK-%s-%d",
			time.Now().Format("20060102"),
			time.Now().UnixNano()%100000000,
		)
	}

	return fmt.Sprintf(
		"JCK-%s-%s",
		time.Now().Format("20060102"),
		strings.ToUpper(hex.EncodeToString(random)),
	)
}
