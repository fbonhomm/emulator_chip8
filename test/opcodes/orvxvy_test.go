package test

import (
	"github.com/fbonhomm/emulator_chip8/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestORVXVY(t *testing.T) {
	var c = cpu.CPU{}

	c.V[2] = 0x0F
	c.V[3] = 0xF0
	c.ORVXVY(0x0320)

	assert.Equal(t, uint8(0xFF), c.V[3])
}
