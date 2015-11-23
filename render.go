package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

func drawAxes(renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.DrawLine(10, 10, 45, 10) // blue line along X axis
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawLine(10, 10, 10, 45) // red line along Y axis
}

func drawRadarBeam(renderer *sdl.Renderer, angle float32) {
	rangle := float64(angle) * 3.14 / 180.0
	// cos(0) = 1
	// sin(0) = 0
	// cx + cos(angle)*100
	// cy + sin(angle)*100
	renderer.SetDrawColor(0, 255, 0, 255)
	cx := winWidth / 2
	cy := winHeight / 2
	renderer.DrawLine(cx, cy, cx+int(math.Cos(rangle)*100.0), cy+int(math.Sin(rangle)*100.0))
}

func render(renderer *sdl.Renderer, angle float32) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	drawAxes(renderer)
	drawRadarBeam(renderer, angle)
	renderer.Present()
}
