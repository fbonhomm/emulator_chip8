package test

import (
	"github.com/fbonhomm/emulator_chip8/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSUBVXVY_VF_1(t *testing.T) {
	var c = cpu.CPU{}

	c.V[8] = 3
	c.V[3] = 2
	c.SUBVXVY(0x0830)

	assert.Equal(t, uint8(1), c.V[8])
	assert.Equal(t, uint8(1), c.V[0xF])
}

func TestSUBVXVY_VF_0(t *testing.T) {
	var c = cpu.CPU{}

	c.V[8] = 2
	c.V[3] = 3
	c.SUBVXVY(0x0830)

	assert.Equal(t, uint8(255), c.V[8])
	assert.Equal(t, uint8(0), c.V[0xF])
}
