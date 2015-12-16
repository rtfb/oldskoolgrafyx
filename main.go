package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ungerik/go3d/vec2"
	"github.com/ungerik/go3d/vec3"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type Camera struct {
	Pos         vec3.T
	Orientation vec3.T // yaw, pitch, roll, in radians. When Orientation is
	// Zero, camera's view vector equals UnitX
	Eye      vec3.T // Eye position behind the screen plane
	FOV      float32
	Viewport vec2.T // Size of the virtual image plane
}

var cam = &Camera{
	Pos:         vec3.Zero,
	Orientation: vec3.Zero,
	Eye:         vec3.Zero,
	FOV:         120,
	Viewport:    vec2.Zero,
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
	cam.Viewport = vec2.T{800, 600}
	cam.Eye[0] = cam.Viewport[0] / 2
	cam.Eye[1] = cam.Viewport[1] / 2
	cam.Eye[2] = float32(-1.0)
	fmt.Printf("cam: %v\n", cam)
	const winTitle = "Go-SDL2 Events"
	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int(cam.Viewport[0]), int(cam.Viewport[1]), sdl.WINDOW_SHOWN)
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
