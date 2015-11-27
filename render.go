package main

import (
	"fmt"

	"github.com/ungerik/go3d/mat3"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type R struct {
	sr      *sdl.Renderer
	bb      *sdl.Surface // backbuffer
	nFrames int          // current frame #, incremented after every frame
	fps     int          // averaged FPS
	font    *ttf.Font
}

func (r *R) drawAxes() {
	r.sr.SetDrawColor(0, 0, 255, 255)
	r.sr.DrawLine(10, 10, 45, 10) // blue line along X axis
	r.sr.SetDrawColor(255, 0, 0, 255)
	r.sr.DrawLine(10, 10, 10, 45) // red line along Y axis
}

func (r *R) drawRadarBeam(angle float32) {
	rangle := angle * 3.14 / 180.0
	m := mat3.Ident
	m.AssignZRotation(rangle)
	view := cam.View
	m.TransformVec3(&view)
	pos2 := cam.Pos
	pos2.Add(&view)
	r.sr.SetDrawColor(0, 255, 0, 255)
	r.sr.DrawLine(int(cam.Pos[0]), int(cam.Pos[1]), int(pos2[0]), int(pos2[1]))
}

func (r *R) drawFPS() {
	fps := fmt.Sprintf("FPS: %d", r.fps)
	surf, err := r.font.RenderUTF8_Blended(fps, sdl.Color{0, 255, 0, 255})
	if err != nil {
		panic(err)
	}
	var srcRect sdl.Rect
	surf.GetClipRect(&srcRect)
	dstRect := sdl.Rect{int32(winWidth) - srcRect.W - 20, 20, srcRect.W, srcRect.H}
	err = surf.Blit(&srcRect, r.bb, &dstRect)
	if err != nil {
		panic(err)
	}
}

func (r *R) doFrame(angle float32) {
	r.sr.SetDrawColor(0, 0, 0, 255)
	r.sr.Clear()
	r.bb.Pixels()[1] = 0xff
	r.drawAxes()
	r.drawRadarBeam(angle)
	r.drawFPS()
	r.sr.Present()
	r.nFrames++
}
