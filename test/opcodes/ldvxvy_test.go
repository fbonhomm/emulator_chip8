package test

import (
	"github.com/fbonhomm/emulator_chip8/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDVXVY(t *testing.T) {
	var c = cpu.CPU{}

	c.V[2] = 15
	c.LDVXVY(0x0320)

	assert.Equal(t, uint8(15), c.V[3])
}
