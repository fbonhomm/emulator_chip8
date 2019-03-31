
package cpu


const SIZE_MEMORY uint16 = 4096 // taille de la memoire total
const START_MEMORY uint16 = 512 // adresse du debut de la memoire
const NB_REGISTER uint16 = 16 // nombre total de registres
const NB_LEVEL_STACK uint16 = 16 // nombre de niveaux pour la stack


type Feature struct {
    memory[SIZE_MEMORY] uint8, // tableaux qui represente la memoire
    register[NB_REGISTER] uint8, //  tableaux qui represente les registres V{0..F}
    stack[NB_LEVEL_STACK] uint16, // tableaux qui represente la stack de sauvegarde
    index_memory uint16, // represente la tete de lecture de la memoire
    index_stack uint8, // represente la tete de lecture de la stack
    system_timer uint8, // mimuterie systeme
    sound_timer uint8, // mimuterie sonore
    address_stock uint16, // stocke une adresse memoire
}

// define getter index_memory
func (f *Feature) get() (uint16) {
  return f.index_memory
}

// decremente les timers
func (f *Feature) decrease() (uint16) {
  if f.system_timer > 0 {
    f.system_timer -= 1
  }
  if f.sound_timer > 0 {
    f.sound_timer -= 1
  }
}
