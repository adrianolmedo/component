package randutil

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUIDStr(t *testing.T) {
	uuids := make(map[string]bool)
	re := regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`)

	for i := 0; i < 1000; i++ {
		uuid := UUIDStr()

		// Comprobar formato
		assert.True(t, re.MatchString(uuid), "UUID con formato incorrecto")

		// Comprobar unicidad
		_, exists := uuids[uuid]
		assert.False(t, exists, "UUID duplicado")
		uuids[uuid] = true
	}
}

func TestUUIDStrWithoutDashes(t *testing.T) {
	uuids := make(map[string]bool)
	re := regexp.MustCompile(`[a-f0-9]{32}`)

	for i := 0; i < 1000; i++ {
		uuid := UUIDStrWithoutDashes()

		// Comprobar formato
		assert.True(t, re.MatchString(uuid), "UUID con formato incorrecto")

		// Comprobar unicidad
		_, exists := uuids[uuid]
		assert.False(t, exists, "UUID duplicado")
		uuids[uuid] = true
	}
}
