package main

import (
	"fmt"
	"os"

	"github.com/ungerik/go3d/vec3"
	"github.com/veandco/go-sdl2/sdl"
)

type Camera struct {
	Pos  vec3.T
	View vec3.T
	FOV  float32
}

var winTitle string = "Go-SDL2 Events"
var winWidth, winHeight int = 800, 600
var angle float32
var cam = &Camera{
	Pos:  vec3.Zero,
	View: vec3.UnitX,
	FOV:  120,
}

func mainLoop(renderer *sdl.Renderer, surface *sdl.Surface, window *sdl.Window) {
	r := &R{renderer, surface}
	quit := make(chan bool)
	for {
		select {
		case <-quit:
			return
		default:
			r.doFrame(angle)
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
	cam.View[0] = 100
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
	mainLoop(renderer, surface, window)
	return 0
}

func main() {
	os.Exit(run())
}
