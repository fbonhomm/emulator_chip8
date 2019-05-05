package test

import (
	"emulator/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDFVX(t *testing.T) {
	var c = cpu.CPU{}

	c.V[2] = 134
	c.LDFVX(0x0200)

	assert.Equal(t, uint16(134*5), c.I)
}
