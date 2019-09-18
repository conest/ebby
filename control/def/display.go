package def

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// DisplayController : 显示控制器，用于监听窗口变化等
type DisplayController struct {
	Win        *pixelgl.Window
	winSize    pixel.Vec
	shareFn    []func(pixel.Vec)
	scenarioFn []func(pixel.Vec)
}

// NewDisplayController : 返回新的显示控制器
func NewDisplayController(w *pixelgl.Window) DisplayController {
	return DisplayController{
		Win:        w,
		winSize:    w.Bounds().Size(),
		shareFn:    make([]func(pixel.Vec), 0),
		scenarioFn: make([]func(pixel.Vec), 0),
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
	for _, fn := range d.shareFn {
		fn(d.winSize)
	}
	for _, fn := range d.scenarioFn {
		fn(d.winSize)
	}
}

// PushShareFn : 添加共享触发窗口变化后的回调函数
func (d *DisplayController) PushShareFn(fn func(pixel.Vec)) {
	d.scenarioFn = append(d.shareFn, fn)
}

// PushScenarioFn : 添加 Scenario 在触发窗口变化后的回调函数
func (d *DisplayController) PushScenarioFn(fn func(pixel.Vec)) {
	d.scenarioFn = append(d.scenarioFn, fn)
}

// ClearScenarioFn : 清除 Scenario 的回调函数
func (d *DisplayController) ClearScenarioFn() {
	d.scenarioFn = []func(pixel.Vec){}
}
