package RNDVXB

import (
	"math/rand"
	"time"
)

// RNDVXB Cxkk - Set Vx = random byte AND kk
func RNDVXB(c *cpu, opcode uint16) {
	var x uint8 = (opcode & 0x0F00) >> 8
	var kk uint8 = (opcode & 0x00FF)

	rand.Seed(time.Now().UnixNano())

	c.register[x] = (rand.Intn(0xFF) & kk)
}
