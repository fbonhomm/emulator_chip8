package test

import (
	"github.com/fbonhomm/emulator_chip8/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJPA(t *testing.T) {
	var c = cpu.CPU{}

	c.JPA(0x0234)

	assert.Equal(t, uint16(0x0232), c.Pc)
}
