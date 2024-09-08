package randutil

import (
	"strings"

	"github.com/google/uuid"
)

// UUIDStr generates a UUID and returns it as a string.
func UUIDStr() string {
	return uuid.New().String()
}

// UUIDStrWithoutDashes generates a UUID and returns it without dashes
// and returns it as a string.
func UUIDStrWithoutDashes() string {
	generated := uuid.New().String()
	return strings.ReplaceAll(generated, "-", "")
}
