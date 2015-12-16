package main

import (
	"fmt"

	"github.com/ungerik/go3d/mat3"
	"github.com/ungerik/go3d/vec2"
	"github.com/ungerik/go3d/vec3"
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

// Project 3D point 'a' from world-space to screen-space
func perspProj(a *vec3.T, cam *Camera) *vec2.T {
	rx := mat3.Ident
	ry := mat3.Ident
	rz := mat3.Ident
	rx.AssignXRotation(-cam.Orientation[0])
	ry.AssignYRotation(-cam.Orientation[1])
	rz.AssignZRotation(-cam.Orientation[2])
	point := vec3.From(a)
	point.Sub(&cam.Pos)
	fmt.Printf("point: %v\n", point)
	temp := mat3.Ident
	temp.AssignMul(&rx, &ry)
	m := mat3.Ident
	m.AssignMul(&temp, &rz)
	m.TransformVec3(&point)
	fmt.Printf("point: %v\n", point)
	sx := cam.Eye[2]/point[2]*point[0] + cam.Eye[0]
	sy := cam.Eye[2]/point[2]*point[1] + cam.Eye[1]
	return &vec2.T{sx, sy}
}

func (r *R) drawRadarBeam() {
	m := mat3.Ident
	m.AssignZRotation(cam.Orientation[0])
	view := vec3.UnitX
	view.Scale(100) // prescale view vector to take meaningful size on screen
	m.TransformVec3(&view)
	pos2 := vec3.From(&cam.Eye)
	pos2[2] = 0
	pos2.Add(&view)
	r.sr.SetDrawColor(0, 255, 0, 255)
	r.sr.DrawLine(int(cam.Eye[0]), int(cam.Eye[1]), int(pos2[0]), int(pos2[1]))
}

func (r *R) drawFPS() {
	fps := fmt.Sprintf("FPS: %d", r.fps)
	surf, err := r.font.RenderUTF8_Blended(fps, sdl.Color{0, 255, 0, 255})
	if err != nil {
		panic(err)
	}
	var srcRect sdl.Rect
	surf.GetClipRect(&srcRect)
	dstRect := sdl.Rect{int32(cam.Viewport[0]) - srcRect.W - 20, 20, srcRect.W, srcRect.H}
	err = surf.Blit(&srcRect, r.bb, &dstRect)
	if err != nil {
		panic(err)
	}
}

func (r *R) doFrame() {
	r.sr.SetDrawColor(0, 0, 0, 255)
	r.sr.Clear()
	r.bb.Pixels()[1] = 0xff
	r.drawAxes()
	r.drawRadarBeam()
	r.drawFPS()
	r.sr.Present()
	r.nFrames++
}
