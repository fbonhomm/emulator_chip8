package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestADDVXVY1(t *testing.T) {
	var c = cpu.CPU{}

	c.V[9] = 0xEE
	c.V[7] = 0x22
	opcodes.ADDVXVY(&c, 0x0970)

	assert.Equal(t, uint8(1), c.V[0xF])
	assert.Equal(t, uint8(0x10), c.V[9])
}

func TestADDVXVY2(t *testing.T) {
	var c = cpu.CPU{}

	c.V[9] = 0xEE
	c.V[7] = 0x11
	opcodes.ADDVXVY(&c, 0x0970)

	assert.Equal(t, uint8(0), c.V[0xF])
	assert.Equal(t, uint8(0xFF), c.V[9])
}
