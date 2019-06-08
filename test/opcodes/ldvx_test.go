package test

import (
	"github.com/fbonhomm/emulator_chip8/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDVX(t *testing.T) {
	var c = cpu.CPU{}

	c.I = 333
	c.Memory[333] = uint8(2)
	c.Memory[334] = uint8(4)
	c.Memory[335] = uint8(6)
	c.LDVX(0x0300)

	assert.Equal(t, uint8(2), c.V[0])
	assert.Equal(t, uint8(4), c.V[1])
	assert.Equal(t, uint8(6), c.V[2])
}
