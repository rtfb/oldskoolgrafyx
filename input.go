package main

import (
	"fmt"
	"math"

	"github.com/ungerik/go3d/vec3"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	oneDegree = math.Pi / 180.0
)

func processInput(quit chan bool) {
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			close(quit)
		case *sdl.MouseMotionEvent:
			fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
				t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
		case *sdl.MouseButtonEvent:
			fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
				t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
		case *sdl.MouseWheelEvent:
			fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
				t.Timestamp, t.Type, t.Which, t.X, t.Y)
		case *sdl.KeyboardEvent:
			if t.Type == sdl.KEYDOWN {
				if t.Keysym.Sym == 1073741904 { // left arrow
					cam.Orientation[0] -= oneDegree
				}
				if t.Keysym.Sym == 1073741903 { // right arrow
					cam.Orientation[0] += oneDegree
				}
				if cam.Orientation[0] < 0 {
					cam.Orientation[0] = 2 * math.Pi
				}
				if cam.Orientation[0] > 2*math.Pi {
					cam.Orientation[0] = 0
				}
				fmt.Printf("cam.Orientation[0] = %f\n", cam.Orientation[0])
				pt := vec3.T{10, 20, 3}
				sp := perspProj(&pt, cam)
				fmt.Printf("sp: %v\n", sp)
			} else if t.Type == sdl.KEYUP {
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tcode:%d\tmodifiers:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Sym, t.Keysym.Mod)
				if t.Keysym.Sym == 27 || t.Keysym.Sym == 1073741881 /* Caps Lock */ {
					close(quit)
				}
			}
		}
	}
}
