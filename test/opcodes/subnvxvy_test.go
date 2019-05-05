package test

import (
	"emulator/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSUBNVXVY_VF_1(t *testing.T) {
	var c = cpu.CPU{}

	c.V[8] = 2
	c.V[3] = 3
	c.SUBNVXVY(0x0830)

	assert.Equal(t, uint8(255), c.V[8])
	assert.Equal(t, uint8(1), c.V[0xF])
}

func TestSUBNVXVY_VF_0(t *testing.T) {
	var c = cpu.CPU{}

	c.V[8] = 3
	c.V[3] = 2
	c.SUBNVXVY(0x0830)

	assert.Equal(t, uint8(1), c.V[8])
	assert.Equal(t, uint8(0), c.V[0xF])
}
