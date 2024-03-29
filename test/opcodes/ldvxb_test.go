package test

import (
	"github.com/fbonhomm/emulator_chip8/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDVXB(t *testing.T) {
	var c = cpu.CPU{}

	c.LDVXB(0x0322)

	assert.Equal(t, uint8(0x22), c.V[3])
}
