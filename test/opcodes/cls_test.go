package test

import (
	"github.com/fbonhomm/emulator_chip8/src/cpu"
	"github.com/fbonhomm/emulator_chip8/src/screen"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLS(t *testing.T) {
	var c = cpu.CPU{}
	var s = screen.SCREEN{}

	s.Pixels[1][1] = 1
	c.CLS(&s)

	assert.Equal(t, uint8(0), s.Pixels[1][1])
}
