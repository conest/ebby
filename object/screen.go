package object

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Screen : 画面，集合了Canvas与Camera功能
type Screen struct {
	Canvas    *pixelgl.Canvas
	pos       pixel.Vec
	Cam       pixel.Matrix
	CamSpeed  float64
	Zoom      float64
	ZoomSpeed float64
	ZoomMin   float64
	ZoomMax   float64
}

// NewScreen : 生成Camera (r: Canvas Rect)
func NewScreen(r pixel.Rect) Screen {
	return Screen{
		Canvas:    pixelgl.NewCanvas(r),
		pos:       pixel.ZV,
		CamSpeed:  300,
		Zoom:      2,
		ZoomSpeed: 1.2,
		ZoomMin:   1,
		ZoomMax:   8,
	}
}

// Update : 根据Pos更新Cam和Canvas位置
func (s *Screen) Update() {
	s.clampPos()
	s.Cam = pixel.IM.Scaled(s.pos, s.Zoom).Moved(s.Canvas.Bounds().Center().Sub(s.pos))
	s.Canvas.SetMatrix(s.Cam)
}

// clampPos: 检查镜头位置防止出Canvas边缘
func (s *Screen) clampPos() {
	posMinX := s.Canvas.Bounds().W() / 2 / s.Zoom
	posMinY := s.Canvas.Bounds().H() / 2 / s.Zoom
	posMaxX := s.Canvas.Bounds().W() - s.Canvas.Bounds().W()/2/s.Zoom
	posMaxY := s.Canvas.Bounds().H() - s.Canvas.Bounds().H()/2/s.Zoom
	s.pos.X = pixel.Clamp(s.pos.X, posMinX, posMaxX)
	s.pos.Y = pixel.Clamp(s.pos.Y, posMinY, posMaxY)
}

// Pos : 返回 Pos
func (s *Screen) Pos() pixel.Vec {
	return s.pos
}

// SetPosCentered : 设置Pos位置为Canvas中央
func (s *Screen) SetPosCentered() {
	s.pos.X = s.Canvas.Bounds().W() / 2
	s.pos.Y = s.Canvas.Bounds().H() / 2
	s.Update()
}

// PosAdd : 修改Pos，xy增加对应的变量乘以镜头速度
func (s *Screen) PosAdd(x, y float64) {
	s.pos.X += s.CamSpeed * x
	s.pos.Y += s.CamSpeed * y
	s.Update()
}

// PosAddX : 修改Pos，x增加对应的变量乘以镜头速度
func (s *Screen) PosAddX(x float64) {
	s.pos.X += s.CamSpeed * x
	s.Update()
}

// PosAddY : 修改Pos，y增加对应的变量乘以镜头速度
func (s *Screen) PosAddY(y float64) {
	s.pos.Y += s.CamSpeed * y
	s.Update()
}

// SetZoom : 设置视角缩放
func (s *Screen) SetZoom(v float64) {
	if v == 0 {
		return
	}
	s.Zoom *= math.Pow(s.ZoomSpeed, v)
	s.Zoom = pixel.Clamp(s.Zoom, s.ZoomMin, s.ZoomMax)
	s.Update()
}

// Unproject : 反向投射，用于映射绝对位置到Canvas相对位置
func (s *Screen) Unproject(v pixel.Vec) pixel.Vec {
	return s.Cam.Unproject(v)
}

// Draw : Draw到目标
func (s *Screen) Draw(target pixel.Target, targetCenter pixel.Vec) {
	s.Canvas.Draw(target, pixel.IM.Moved(targetCenter))
}
