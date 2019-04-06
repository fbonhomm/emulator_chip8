package cpu

// 0x000-0x1FF - Chip 8 interpreter (contains font set in emu)
// 0x050-0x0A0 - Used for the built in 4x5 pixel font set (0-F)
// 0x200-0xFFF - Program ROM and work RAM

const sizeMemory uint16 = 4096 // taille de la memoire total
const startMemory uint16 = 512 // adresse du debut de la memoire
const nbrRegister uint8 = 16   // nombre total de registres
const nbrLevelStack uint8 = 16 // nombre de niveaux pour la stack
const nbrOpcode uint8 = 35     // nombre de niveaux pour la stack
// const OPCODE [] uint16 {
//   0x00E0, 0x00EE, 0x0NNN, 0x1NNN, 0x2NNN, 0x3XKK, 0x4XKK, 0x5XY0, 0x6XKK,
//   0x7XKK, 0x8XY0, 0x8XY1, 0x8XY2, 0x8XY3, 0x8XY4, 0x8XY5, 0x8XY6, 0x8XY7,
//   0x8XYE, 0x9XY0, 0xANNN, 0xBNNN, 0xCXKK, 0xDXYN, 0xEX9E, 0xEXA1, 0xFX07,
//   0xFX0A, 0xFX15, 0xFX18, 0xFX1E, 0xFX29, 0xFX33, 0xFX55, 0xFX65,
// }
const opcodeMask = []uint16{
	0x00F0, 0x00FF, 0x0000, 0xF000, 0xF000, 0xF000, 0xF000, 0xF00F, 0xF000,
	0xF000, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F,
	0xF00F, 0xF00F, 0xF000, 0xF000, 0xF000, 0xF000, 0xF0FF, 0xF0FF, 0xF0FF,
	0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF,
}
const opcodeID = []uint16{
	0x00E0, 0x00EE, 0x0000, 0x1000, 0x2000, 0x3000, 0x4000, 0x5000, 0x6000,
	0x7000, 0x8000, 0x8001, 0x8002, 0x8003, 0x8004, 0x8005, 0x8006, 0x8007,
	0x800E, 0x9000, 0xA000, 0xB000, 0xC000, 0xD000, 0xE09E, 0xE0A1, 0xF007,
	0xF00A, 0xF015, 0xF018, 0xF01E, 0xF029, 0xF033, 0xF055, 0xF065,
}

type feature struct {
	memory       [sizeMemory]uint8     // tableaux qui represente la memoire
	register     [nbrRegister]uint8    //  tableaux qui represente les registres V{0..F}
	stack        [nbrLevelStack]uint16 // tableaux qui represente la stack de sauvegarde
	indexMemory  uint16                // represente la tete de lecture de la memoire
	indexStack   uint8                 // represente la tete de lecture de la stack
	systemTimer  uint8                 // mimuterie systeme
	soundTimer   uint8                 // mimuterie sonore
	addressStock uint16                // stocke une adresse memoire
}

// define getter index_memory
func (f *feature) get() uint16 {
	return f.index_memory
}

// return index opcode or -1 if not found
func (f *feature) identifyOpcode(address uint16) uint8 {
	for idx := 0; idx < nbrOpcode; idx++ {
		if address&opcodeMask[idx] == opcodeID[idx] {
			return idx
		}
	}

	return -1
}

// decremente les timers
func (f *feature) decrease() uint16 {
	if f.systemTimer > 0 {
		f.systemTimer--
	}
	if f.soundTimer > 0 {
		f.soundTimer--
	}
}
