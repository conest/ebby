package scenario

import (
	"ebby/control/strdef"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Run : 主运行
func (s *Scenario) Run(w *pixelgl.Window) strdef.Request {

	req := strdef.DefaultRequest
	last := time.Now()

	// DEBUG: debug mode
	// fps
	fps := strdef.NewFps(w, s.sdata.Resource.DebugAtlas)

	// 执行循环
	for {
		if w.Closed() {
			req = strdef.Request{Terminate: true}
			return req
		}

		// Delta Time
		dt := time.Since(last).Seconds()
		last = time.Now()

		s.inputHandle(w, dt)

		if r := s.excute(dt); !r.Continue {
			req = r
			return req
		}

		w.Clear(colornames.Black)
		s.draw(w)
		fps.Update() // DEBUG: debug mode
		w.Update()

		<-s.sTicker.C
	}
}

// excute : 数据执行
func (s *Scenario) excute(dt float64) strdef.Request {
	select {
	case <-s.eTicker.C:
		r := s.ins.Excuter(dt)
		return r
	default:
		return strdef.Request{Continue: true}
	}
}

// draw : 绘图
func (s *Scenario) draw(w *pixelgl.Window) {
	s.ins.Drawer(w)
	// DEBUG: debug mode
	s.sdata.Tool.DebugLogger.Draw(w, pixel.IM)
}

// inputHandle : 输入监听
func (s *Scenario) inputHandle(w *pixelgl.Window, dt float64) {
	s.ins.InputHandle(w, dt)
}
