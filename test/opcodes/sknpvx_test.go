package test

import (
	"github.com/fbonhomm/emulator_chip8/src/cpu"
	"github.com/fbonhomm/emulator_chip8/src/keyboard"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSKNPVX_PC_2(t *testing.T) {
	var c = cpu.CPU{}

	key := keyboard.GetKey(8)
	key.Pressed = 0
	c.SKNPVX(0x0800)

	assert.Equal(t, uint16(2), c.Pc)
}

func TestSKNPVX_PC_0(t *testing.T) {
	var c = cpu.CPU{}

	key := keyboard.GetKey(8)
	key.Pressed = 1
	c.SKNPVX(0x0800)

	assert.Equal(t, uint16(0), c.Pc)
}
