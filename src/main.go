package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)


func run() {
  cfg := pixelgl.WindowConfig{
    Title: "Test",
    Bounds: pixel.R(0, 0, 1024, 768),
  }

  win, err := pixelgl.NewWindow(cfg)

  if (err != nil) {
    panic(err)
  }

  for !win.Closed() {
    win.Update()
  }
}

func main() {
	// var test[10] uint8

	// fmt.Printf("%v\n", test)
  // pixelgl.Run(run)
	fmt.Printf("%v\n", 0x00E0)
  // fmt.Printf("%v\n", 0x00EE)
  // fmt.Printf("%v\n", 0x0NNN)
  // fmt.Printf("%v\n", 0x1NNN)
  // fmt.Printf("%v\n", 0x2NNN)
  // fmt.Printf("%v\n", 0x3XKK)
  // fmt.Printf("%v\n", 0x4XKK)
  // fmt.Printf("%v\n", 0x5XY0)
  // fmt.Printf("%v\n", 0x6XKK)
  // fmt.Printf("%v\n", 0x7XKK)
  // fmt.Printf("%v\n", 0x8XY0)
  // fmt.Printf("%v\n", 0x8XY1)
  // fmt.Printf("%v\n", 0x8XY2)
  // fmt.Printf("%v\n", 0x8XY3)
  // fmt.Printf("%v\n", 0x8XY4)
  // fmt.Printf("%v\n", 0x8XY5)
  // fmt.Printf("%v\n", 0x8XY6)
  // fmt.Printf("%v\n", 0x8XY7)
  // fmt.Printf("%v\n", 0x8XYE)
  // fmt.Printf("%v\n", 0x9XY0)
  // fmt.Printf("%v\n", 0xANNN)
  // fmt.Printf("%v\n", 0xBNNN)
  // fmt.Printf("%v\n", 0xCXKK)
  // fmt.Printf("%v\n", 0xDXYN)
  // fmt.Printf("%v\n", 0xEX9E)
  // fmt.Printf("%v\n", 0xEXA1)
  // fmt.Printf("%v\n", 0xFX07)
  // fmt.Printf("%v\n", 0xFX0A)
  // fmt.Printf("%v\n", 0xFX15)
  // fmt.Printf("%v\n", 0xFX18)
  // fmt.Printf("%v\n", 0xFX1E)
  // fmt.Printf("%v\n", 0xFX29)
  // fmt.Printf("%v\n", 0xFX33)
  // fmt.Printf("%v\n", 0xFX55)
  // fmt.Printf("%v\n", 0xFX65)
}
