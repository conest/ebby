package scene

import (
	"github.com/conest/ebby/game/def"
	"github.com/conest/ebby/game/tool"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Run : 主运行
func (s *Scene) Run(w *pixelgl.Window) def.Request {

	req := def.DefaultRequest
	dts := NewDeltaTimer()

	// DEBUG: debug mode
	// fps
	fps := tool.NewFps(w, s.sdata.Resource.DebugAtlas)

	// 执行循环
	for {
		if window.Closed() {
			req = def.Request{Terminate: true}
			return req
		}

		// Delta Time
		deltaTime := dts.GetDeltaTime()

		scenario.sdata.Tool.Display.Update()
		scenario.inputHandle(window, deltaTime)

		if r := scenario.excute(dts); !r.Continue {
			req = r
			return req
		}

		// Update Frame
		window.Clear(colornames.Black)
		scenario.draw(window, deltaTime)
		fps.Update() // DEBUG: debug mode
		window.Update()

		<-scenario.sTicker.C
	}
}

// excute : 数据执行
func (s *Scene) excute(dts DeltaTime) def.Request {
	select {
	case <-scenario.eTicker.C:
		r := scenario.ins.Excuter(dtr)
		return r
	default:
		return def.Request{Continue: true}
	}
}

// draw : 绘图
func (s *Scene) draw(w *pixelgl.Window, dt float64) {
	s.ins.Drawer(w, dt)
	// DEBUG: debug mode
	scenario.sdata.Tool.DebugLogger.Draw(window, pixel.IM)
}

// inputHandle : 输入监听
func (s *Scene) inputHandle(w *pixelgl.Window, dt float64) {
	s.ins.InputHandle(w, dt)
}
