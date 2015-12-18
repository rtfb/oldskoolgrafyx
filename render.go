package main

import (
	"fmt"

	"github.com/ungerik/go3d/mat3"
	"github.com/ungerik/go3d/mat4"
	"github.com/ungerik/go3d/vec2"
	"github.com/ungerik/go3d/vec3"
	"github.com/ungerik/go3d/vec4"
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

func printM4(m *mat4.T) {
	for i := 0; i < 4; i++ {
		fmt.Printf("[%f %f %f %f]\n", m[0][i], m[1][i], m[2][i], m[3][i])
	}
}

func printM3(m *mat3.T) {
	for i := 0; i < 3; i++ {
		fmt.Printf("[%f %f %f]\n", m[0][i], m[1][i], m[2][i])
	}
}

func mkExtrinsicCameraMtx(cam *Camera) *mat4.T {
	mRot := mat3.Ident
	mRot.AssignEulerRotation(cam.Orientation[0], cam.Orientation[1], cam.Orientation[2])
	mRot.Transpose()
	printM3(&mRot)
	m := mat4.From(&mRot)
	mRot.Scale(-1)
	camPos := vec3.From(&cam.Pos)
	mRot.TransformVec3(&camPos)
	m[3] = vec4.From(&camPos)
	return &m
}

// Project 3D point 'a' from world-space to screen-space
func perspProj(a *vec3.T, cam *Camera) *vec2.T {
	m := mkExtrinsicCameraMtx(cam)
	printM4(m)
	sp := vec3.From(a)
	m.TransformVec3(&sp)
	return &vec2.T{sp[0], sp[1]}
}

func (r *R) drawRadarBeam() {
	m := mat3.Ident
	m.AssignZRotation(cam.Orientation[0])
	view := vec3.UnitX
	view.Scale(100) // prescale view vector to take meaningful size on screen
	m.TransformVec3(&view)
	pos2 := vec3.From(&cam.Pos)
	pos2[2] = 0
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
