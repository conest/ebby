package scene

import (
	"ebby/game/def"
	"ebby/game/tool"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Run : 主运行
func (s *Scene) Run(w *pixelgl.Window) def.Request {

	req := def.DefaultRequest
	dts := NewDT()

	// DEBUG: debug mode
	// fps
	fps := tool.NewFps(w, s.sdata.Resource.DebugAtlas)

	// 执行循环
	for {
		if w.Closed() {
			req = def.Request{Terminate: true}
			return req
		}

		// Delta Time
		dt := dts.Get()

		s.sdata.Tool.Display.Update()
		s.inputHandle(w, dt)

		if r := s.excute(dts); !r.Continue {
			req = r
			return req
		}

		w.Clear(colornames.Black)
		s.draw(w, dt)
		fps.Update() // DEBUG: debug mode
		w.Update()

		<-s.sTicker.C
	}
}

// excute : 数据执行
func (s *Scene) excute(dts DeltaTime) def.Request {
	select {
	case <-s.eTicker.C:
		r := s.ins.Excuter(dts)
		return r
	default:
		return def.Request{Continue: true}
	}
}

// draw : 绘图
func (s *Scene) draw(w *pixelgl.Window, dt float64) {
	s.ins.Drawer(w, dt)
	// DEBUG: debug mode
	s.sdata.Tool.DebugLogger.Draw(w, pixel.IM)
}

// inputHandle : 输入监听
func (s *Scene) inputHandle(w *pixelgl.Window, dt float64) {
	s.ins.InputHandle(w, dt)
}
