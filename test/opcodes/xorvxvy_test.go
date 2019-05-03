package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXORVXVY(t *testing.T) {
	var c = cpu.CPU{}

	c.V[2] = 0x0F
	c.V[3] = 0xF1
	opcodes.XORVXVY(&c, 0x0320)

	assert.Equal(t, uint8(0xFE), c.V[3])
}
