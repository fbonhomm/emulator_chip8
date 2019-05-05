package test

import (
	"emulator/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJPV0(t *testing.T) {
	var c = cpu.CPU{}

	c.V[0x0] = 0x2
	c.JPV0(0x0234)

	assert.Equal(t, uint16(0x0234), c.Pc)
}
