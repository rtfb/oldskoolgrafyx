package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type R struct {
	sr *sdl.Renderer
}

func (r *R) drawAxes() {
	r.sr.SetDrawColor(0, 0, 255, 255)
	r.sr.DrawLine(10, 10, 45, 10) // blue line along X axis
	r.sr.SetDrawColor(255, 0, 0, 255)
	r.sr.DrawLine(10, 10, 10, 45) // red line along Y axis
}

func (r *R) drawRadarBeam(angle float32) {
	rangle := float64(angle) * 3.14 / 180.0
	// cos(0) = 1
	// sin(0) = 0
	// cx + cos(angle)*100
	// cy + sin(angle)*100
	r.sr.SetDrawColor(0, 255, 0, 255)
	cx := winWidth / 2
	cy := winHeight / 2
	r.sr.DrawLine(cx, cy, cx+int(math.Cos(rangle)*100.0), cy+int(math.Sin(rangle)*100.0))
}

func (r *R) doFrame(angle float32) {
	r.sr.SetDrawColor(0, 0, 0, 255)
	r.sr.Clear()
	r.drawAxes()
	r.drawRadarBeam(angle)
	r.sr.Present()
}
