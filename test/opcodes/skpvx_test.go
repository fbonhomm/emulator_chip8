package test

import (
	"emulator/src/cpu"
	"emulator/src/keyboard"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSKPVX_PC_2(t *testing.T) {
	var c = cpu.CPU{}

	key := keyboard.GetKey(8)
	key.Pressed = 1
	c.SKPVX(0x0800)

	assert.Equal(t, uint16(2), c.Pc)
}

func TestSKPVX_PC_0(t *testing.T) {
	var c = cpu.CPU{}

	key := keyboard.GetKey(8)
	key.Pressed = 0
	c.SKPVX(0x0800)

	assert.Equal(t, uint16(0), c.Pc)
}
