package test

import (
	"emulator/src/cpu"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDVXDT(t *testing.T) {
	var c = cpu.CPU{}

	c.DT = 23
	c.LDVXDT(0x0300)

	assert.Equal(t, uint8(23), c.V[3])
}
