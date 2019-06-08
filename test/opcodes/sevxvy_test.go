package test

import (
	"github.com/fbonhomm/emulator_chip8/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSEVXVY_SET_PC(t *testing.T) {
	var c = cpu.CPU{}

	c.V[0x3] = 0xC
	c.V[0xE] = 0xC
	c.SEVXVY(0x03E0)

	assert.NotEqual(t, uint8(1), c.Pc)
}

func TestSEVXVY_NOT_SET_PC(t *testing.T) {
	var c = cpu.CPU{}

	c.V[0x3] = 0xC
	c.V[0xE] = 0xA
	c.SEVXVY(0x03E0)

	assert.NotEqual(t, uint8(0), c.Pc)
}

//
