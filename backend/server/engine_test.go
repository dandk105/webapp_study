package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEngine(t *testing.T) {
	address := ":5001"
	engine := CreateEngine()

	assert.Equal(t, engine.Addr, address)
}
