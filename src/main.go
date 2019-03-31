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
	var test[10] uint8

	fmt.Printf("%v\n", test)
  // pixelgl.Run(run)
}
