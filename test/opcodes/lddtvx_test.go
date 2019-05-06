package test

import (
	"emulator/src/cpu"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDDTVX(t *testing.T) {
	var c = cpu.CPU{}

	c.V[2] = 134
	c.LDDTVX(0x0200)

	assert.Equal(t, uint8(134), c.DT)
}
