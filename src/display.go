package display

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const PIXEL_SIZE unit8 = 8
const PIXEL_BLACK unit8 = 0
const PIXEL_WHITE unit8 = 1
const HEIGHT unit8 = 32
const WIDTH unit8 = 64
const WINDOW_HEIGHT unit16 = HEIGHT * PIXEL_SIZE
const WINDOW_WIDTH unit16 = WIDTH * PIXEL_SIZE


type Pixel struct {
  color uint8, // couleur du pixel
}

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
