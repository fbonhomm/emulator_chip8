package test

import (
	"emulator/src/cpu"
	"emulator/src/screen"
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
