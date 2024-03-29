package test

import (
	"github.com/fbonhomm/emulator_chip8/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRNDVXB(t *testing.T) {
	var c = cpu.CPU{}

	c.RNDVXB(0x03EC)

	assert.NotEqual(t, uint8(0), c.V[3])
}
