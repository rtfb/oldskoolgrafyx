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

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var event sdl.Event
	var running bool
	var err error
	var angle float32
	cam := Camera{
		Pos:  vec3.Zero,
		View: vec3.UnitX,
		FOV:  120,
	}
	fmt.Printf("cam: %v\n", cam)
	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	running = true
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
			case *sdl.MouseWheelEvent:
				fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y)
			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == 1073741904 { // left arrow
					angle -= 1.0
				}
				if t.Keysym.Sym == 1073741903 { // right arrow
					angle += 1.0
				}
				if angle < 0 {
					angle = 360
				}
				if angle > 360 {
					angle = 0
				}
				fmt.Printf("angle = %f\n", angle)
			case *sdl.KeyUpEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tcode:%d\tmodifiers:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Sym, t.Keysym.Mod)
				if t.Keysym.Sym == 27 || t.Keysym.Sym == 1073741881 /* Caps Lock */ {
					running = false
				}
			}
		}
		render(renderer, angle)
	}

	return 0
}

func main() {
	os.Exit(run())
}
