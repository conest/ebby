package object

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Screen : 画面，集合了Canvas与Camera功能
type Screen struct {
	Canvas     *pixelgl.Canvas
	ScreenSize pixel.Vec
	pos        pixel.Vec
	cam        pixel.Matrix
	CamSpeed   float64
	Zoom       float64
	ZoomSpeed  float64
	ZoomMin    float64
	ZoomMax    float64
	IntZoom    bool
}

// NewScreen : 生成Camera (r: Canvas Rect)
func NewScreen(r pixel.Rect, size pixel.Vec) Screen {
	return Screen{
		Canvas:     pixelgl.NewCanvas(r),
		ScreenSize: size,
		pos:        pixel.ZV,
		CamSpeed:   300,
		Zoom:       1,
		ZoomSpeed:  1.2,
		ZoomMin:    1,
		ZoomMax:    8,
		IntZoom:    false,
	}
}

// Update : 根据Pos更新Cam和Canvas位置
func (s *Screen) Update() {
	s.clampPos()
	s.cam = pixel.IM.Scaled(s.pos, s.Zoom).Moved(s.Canvas.Bounds().Center().Sub(s.pos))
	s.Canvas.SetMatrix(s.cam)
}

// clampPos: 检查镜头位置防止出Canvas边缘
func (s *Screen) clampPos() {
	posMinX := s.Canvas.Bounds().W() / 2 / s.Zoom
	posMinY := s.Canvas.Bounds().H() / 2 / s.Zoom
	posMaxX := s.ScreenSize.X - s.Canvas.Bounds().W()/2/s.Zoom
	posMaxY := s.ScreenSize.Y - s.Canvas.Bounds().H()/2/s.Zoom
	s.pos.X = pixel.Clamp(s.pos.X, posMinX, posMaxX)
	s.pos.Y = pixel.Clamp(s.pos.Y, posMinY, posMaxY)
}

// limitCam: 限制Zoom大小防止出Canvas边缘
func (s *Screen) limitZoom() {
	widthZoomLimit := s.Canvas.Bounds().W() / s.ScreenSize.X
	heightZoomLimit := s.Canvas.Bounds().H() / s.ScreenSize.Y
	ZoomLimit := math.Max(widthZoomLimit, heightZoomLimit)
	if s.Zoom < ZoomLimit {
		s.Zoom = ZoomLimit
	}
}

// Pos : 返回 Pos
func (s *Screen) Pos() pixel.Vec {
	return s.pos
}

// SetPos : 设置Pos位置
func (s *Screen) SetPos(v pixel.Vec) {
	s.pos = v
	s.Update()
}

// SetPosCentered : 设置Pos位置为Canvas中央
func (s *Screen) SetPosCentered() {
	s.pos.X = s.ScreenSize.X / 2
	s.pos.Y = s.ScreenSize.Y / 2
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

// Cam : 返回 Cam
func (s *Screen) Cam() pixel.Matrix {
	return s.cam
}

// SetZoom : 设置视角缩放
func (s *Screen) SetZoom(v float64) {
	s.Zoom = v
	if s.IntZoom {
		s.Zoom = math.Floor(s.Zoom)
	}
	s.Zoom = pixel.Clamp(s.Zoom, s.ZoomMin, s.ZoomMax)
	s.limitZoom()
	s.Update()
}

// ScrollZoom : 滚动视角缩放(不推荐在IntZoom开启时使用)
func (s *Screen) ScrollZoom(v float64) {
	if v == 0 {
		return
	}
	s.Zoom *= math.Pow(s.ZoomSpeed, v)
	if s.IntZoom {
		s.Zoom = math.Floor(s.Zoom)
	}
	s.Zoom = pixel.Clamp(s.Zoom, s.ZoomMin, s.ZoomMax)
	s.limitZoom()
	s.Update()
}

// ScrollZoomSteped : 步进式滚动视角缩放(推荐在IntZoom开启时使用)
func (s *Screen) ScrollZoomSteped(v float64) {
	if v == 0 {
		return
	}
	if v > 0 {
		s.Zoom++
	} else {
		s.Zoom--
	}
	if s.IntZoom {
		s.Zoom = math.Floor(s.Zoom)
	}
	s.Zoom = pixel.Clamp(s.Zoom, s.ZoomMin, s.ZoomMax)
	s.limitZoom()
	s.Update()
}

// Unproject : 反向投射，用于映射绝对位置到Canvas相对位置
func (s *Screen) Unproject(v pixel.Vec) pixel.Vec {
	return s.cam.Unproject(v)
}

// Draw : Draw到目标
func (s *Screen) Draw(target pixel.Target, targetCenter pixel.Vec) {
	s.Canvas.Draw(target, pixel.IM.Moved(targetCenter))
}

// DisplayCallBack : Screen触发窗口变化后的回调函数
func (s *Screen) DisplayCallBack(winVec pixel.Vec) {
	s.Canvas.SetBounds(pixel.R(0, 0, winVec.X, winVec.Y))
	s.limitZoom()
	s.Update()
}
