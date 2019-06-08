package test

import (
	"github.com/fbonhomm/emulator_chip8/src/cpu"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestADDIVX(t *testing.T) {
	var c = cpu.CPU{}

	c.V[2] = 2
	c.ADDIVX(0x0200)

	assert.Equal(t, uint16(2), c.I)
}
