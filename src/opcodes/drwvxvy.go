package DRWVXVY

// DRWVXVY Dxyn - Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
func DRWVXVY(c *cpu, opcode uint16, s *screen) {
	var x uint8 = uint8((opcode & 0x0F00) >> 8)
	var y uint8 = uint8((opcode & 0x00F0) >> 4)
	var n uint16 = uint16((opcode & 0x000F))
	var pixel uint8

	c.register[0xF] = 0
	for i := uint16(0); i < n; i++ {
		pixel = (c.memory[c.I+i] & 0x00FF)

		for j := 0; j < 8; j++ {
			if (pixel & (0x80 >> j)) != 0 {
				if s.pixels[y+i][x+j] == 1 {
					c.register[0xF] = 1
				}
				s.pixels[y+i][x+j] ^= 1
			}
		}

		if (i + 1) < n {
			pixel = (c.memory[c.I+i] & 0xFF00) >> 8

			i++
			for j := 0; j < 8; j++ {
				if (pixel & (0x80 >> j)) != 0 {
					if s.pixels[y+i][x+j] == 1 {
						c.register[0xF] = 1
					}
					s.pixels[y+i][x+j] ^= 1
				}
			}
		}
	}
}
