package test

import (
	"github.com/fbonhomm/emulator_chip8/src/cpu"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRET(t *testing.T) {
	var c = cpu.CPU{}

	c.Sp = 2
	c.Stack[2] = 432
	c.RET()

	assert.Equal(t, uint16(432), c.Pc)
}
