package cpu

import (
	"math/rand"
	"time"
)

// RNDVXB Cxkk - Set Vx = random byte AND kk
func (c *CPU) RNDVXB(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)
	kk := uint8((opcode & 0x00FF))

	rand.Seed(time.Now().UnixNano())

	c.V[x] = uint8(rand.Intn(0xFF)) & kk
}
