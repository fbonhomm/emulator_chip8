package test

import (
	"emulator/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSEVXB_SET_PC(t *testing.T) {
	var c = cpu.CPU{}

	c.V[3] = 0xEC
	c.SEVXB(0x03EC)

	assert.NotEqual(t, uint8(1), c.Pc)
}

func TestSEVXB_NOT_SET_PC(t *testing.T) {
	var c = cpu.CPU{}

	c.V[3] = 0xEB
	c.SEVXB(0x03EC)

	assert.NotEqual(t, uint8(0), c.Pc)
}
