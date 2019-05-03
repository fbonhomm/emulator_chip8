package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestANDVXVY(t *testing.T) {
	var c = cpu.CPU{}

	c.V[9] = 0xF0
	c.V[7] = 0x0F
	opcodes.ANDVXVY(&c, 0x0970)

	assert.Equal(t, uint8(0x00), c.V[9])
}
