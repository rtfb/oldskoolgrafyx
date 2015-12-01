package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ungerik/go3d/vec3"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type Camera struct {
	Pos         vec3.T
	Orientation vec3.T // yaw, pitch, roll, in radians. When Orientation is
	// Zero, camera's view vector equals UnitX
	FOV float32
}

var winTitle string = "Go-SDL2 Events"
var winWidth, winHeight int = 800, 600
var cam = &Camera{
	Pos:         vec3.Zero,
	Orientation: vec3.Zero,
	FOV:         120,
}

func mainLoop(renderer *sdl.Renderer, surface *sdl.Surface, window *sdl.Window) {
	font, err := ttf.OpenFont("fonts/sfd/FreeMono.ttf", 16)
	if err != nil {
		panic(err)
	}
	r := &R{
		sr:      renderer,
		bb:      surface,
		nFrames: 0,
		fps:     0,
		font:    font,
	}
	quit := make(chan bool)
	fpsTicker := time.NewTicker(1 * time.Second)
	fpsQuit := make(chan struct{})
	for {
		select {
		case <-quit:
			close(fpsQuit)
			return
		case <-fpsTicker.C:
			r.fps = r.nFrames
			r.nFrames = 0
		default:
			r.doFrame()
			processInput(quit)
		}
		window.UpdateSurface()
		sdl.Delay(1) // yield CPU
	}
}

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error
	cam.Pos[0] = float32(winWidth / 2)
	cam.Pos[1] = float32(winHeight / 2)
	fmt.Printf("cam: %v\n", cam)
	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	renderer, err = sdl.CreateSoftwareRenderer(surface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()
	err = ttf.Init()
	if err != nil {
		panic(err)
	}
	mainLoop(renderer, surface, window)
	return 0
}

func main() {
	os.Exit(run())
}
