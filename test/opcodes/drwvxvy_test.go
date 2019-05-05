package test

import (
	"emulator/src/cpu"
	
	"emulator/src/screen"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDRWVXVY_WITHOUT_COLLISION(t *testing.T) {
	var c = cpu.CPU{}
	var s = screen.SCREEN{}

	c.I = 3000
	c.Memory[2999] = 0x3C // 00111100
	c.Memory[3000] = 0x55 // 01010101
	c.Memory[3001] = 0xFF // 11111111
	c.Memory[3002] = 0xAA // 10101010
	c.Memory[3003] = 0x3C // 00111100

	c.DRWVXVY(&s, 0x0133)

	assert.Equal(t, uint8(0), c.V[0xF])

	assert.Equal(t, []uint8{0, 0, 0, 0, 0, 0, 0, 0}, s.Pixels[2][1:(1+8)])
	assert.Equal(t, []uint8{0, 1, 0, 1, 0, 1, 0, 1}, s.Pixels[3][1:(1+8)])
	assert.Equal(t, []uint8{1, 1, 1, 1, 1, 1, 1, 1}, s.Pixels[4][1:(1+8)])
	assert.Equal(t, []uint8{1, 0, 1, 0, 1, 0, 1, 0}, s.Pixels[5][1:(1+8)])
	assert.Equal(t, []uint8{0, 0, 0, 0, 0, 0, 0, 0}, s.Pixels[6][1:(1+8)])
}

func TestDRWVXVY_WITH_COLLISION(t *testing.T) {
	var c = cpu.CPU{}
	var s = screen.SCREEN{}

	c.I = 3000
	s.Pixels[4][1] = 1    // collision with c.Memory[3001]
	c.Memory[2999] = 0x3C // 00111100
	c.Memory[3000] = 0x55 // 01010101
	c.Memory[3001] = 0xFF // 11111111
	c.Memory[3002] = 0xAA // 10101010
	c.Memory[3003] = 0x3C // 00111100

	c.DRWVXVY(&s, 0x0133)

	assert.Equal(t, uint8(1), c.V[0xF])

	assert.Equal(t, []uint8{0, 0, 0, 0, 0, 0, 0, 0}, s.Pixels[2][1:(1+8)])
	assert.Equal(t, []uint8{0, 1, 0, 1, 0, 1, 0, 1}, s.Pixels[3][1:(1+8)])
	assert.Equal(t, []uint8{0, 1, 1, 1, 1, 1, 1, 1}, s.Pixels[4][1:(1+8)]) // first element erase
	assert.Equal(t, []uint8{1, 0, 1, 0, 1, 0, 1, 0}, s.Pixels[5][1:(1+8)])
	assert.Equal(t, []uint8{0, 0, 0, 0, 0, 0, 0, 0}, s.Pixels[6][1:(1+8)])
}
