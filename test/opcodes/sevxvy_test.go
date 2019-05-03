package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSEVXVY_SET_PC(t *testing.T) {
	var c = cpu.CPU{}

	c.V[0x3] = 0xC
	c.V[0xE] = 0xC
	opcodes.SEVXVY(&c, 0x03E0)

	assert.NotEqual(t, uint8(1), c.Pc)
}

func TestSEVXVY_NOT_SET_PC(t *testing.T) {
	var c = cpu.CPU{}

	c.V[0x3] = 0xC
	c.V[0xE] = 0xA
	opcodes.SEVXVY(&c, 0x03E0)

	assert.NotEqual(t, uint8(0), c.Pc)
}

//
