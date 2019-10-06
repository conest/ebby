package sys

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// DisplayController : 显示控制器，用于监听窗口变化等
type DisplayController struct {
	Win      *pixelgl.Window
	winSize  pixel.Vec
	publicFn []func(pixel.Vec)
	sceneFn  []func(pixel.Vec)
}

// NewDisplayController : 返回新的显示控制器
func NewDisplayController(win *pixelgl.Window) DisplayController {
	return DisplayController{
		Win:      win,
		winSize:  win.Bounds().Size(),
		publicFn: make([]func(pixel.Vec), 0),
		sceneFn:  make([]func(pixel.Vec), 0),
	}
}

// SetWindowSize : 设定窗口大小并自动居中
func (d *DisplayController) SetWindowSize(w, h float64) {
	d.Win.SetBounds(pixel.R(0, 0, w, h))
	mw, mh := pixelgl.PrimaryMonitor().Size()
	if mw < w || mh < h {
		d.Win.SetPos(pixel.V(0, 0))
		return
	}
	d.Win.SetPos(pixel.V((mw-w)/2, (mh-h)/2))
}

// Update : 监听并更新Window变化
func (d *DisplayController) Update() {
	if !d.winSize.Eq(d.Win.Bounds().Size()) {
		d.winSize = d.Win.Bounds().Size()
		d.callBack()
	}
}

func (d *DisplayController) callBack() {
	for _, fn := range d.publicFn {
		fn(d.winSize)
	}
	for _, fn := range d.sceneFn {
		fn(d.winSize)
	}
}

// PushPublicFn : 添加共享触发窗口变化后的回调函数
func (d *DisplayController) PushPublicFn(fn func(pixel.Vec)) {
	d.sceneFn = append(d.publicFn, fn)
}

// PushSceneFn : 添加 Scene 在触发窗口变化后的回调函数
func (d *DisplayController) PushSceneFn(fn func(pixel.Vec)) {
	d.sceneFn = append(d.sceneFn, fn)
}

// ClearSceneFn : 清除 Scene 的回调函数
func (d *DisplayController) ClearSceneFn() {
	d.sceneFn = []func(pixel.Vec){}
}
