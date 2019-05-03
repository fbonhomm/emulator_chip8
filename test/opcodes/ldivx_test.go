package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDIVX(t *testing.T) {
	var c = cpu.CPU{}

	c.I = 333
	c.V[0] = uint8(2)
	c.V[1] = uint8(4)
	c.V[2] = uint8(6)
	opcodes.LDIVX(&c, 0x0300)

	assert.Equal(t, uint8(2), c.Memory[333])
	assert.Equal(t, uint8(4), c.Memory[334])
	assert.Equal(t, uint8(6), c.Memory[335])
}
