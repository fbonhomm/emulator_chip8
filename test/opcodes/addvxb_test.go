package test

import (
	"emulator/src/cpu"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestADDVXB(t *testing.T) {
	var c = cpu.CPU{}

	c.V[2] = 1
	c.ADDVXB(0x0233)

	assert.Equal(t, uint8(0x34), c.V[2])
}
