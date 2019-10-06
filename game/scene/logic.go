package scene

import (
	"github.com/conest/ebby/game/def"
	"github.com/conest/ebby/game/tool"

	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

// Run : 主运行
func (s *Scene) Run() def.Request {

	// DEBUG: debug mode
	// fps
	win := s.gamedata.Sys.Win
	fps := tool.NewFps(win, s.gamedata.Tool.DebugAtlas)

	// 执行循环
	for {
		if win.Closed() {
			s.req = def.Request{Terminate: true}
			return s.req
		}
		s.gamedata.Sys.Display.Update()

		// Delta Time
		s.dti.Update()
		s.inputHandle()

		if r := s.excute(); !r.Continue {
			s.req = r
			return s.req
		}

		win.Clear(colornames.Black)
		s.draw()
		fps.Update() // DEBUG: debug mode
		win.Update()

		<-s.sTicker.C
	}
}

// excute : 数据执行
func (s *Scene) excute() def.Request {
	return s.ins.Excuter(s.dti.Dt)
}

// draw : 绘图
func (s *Scene) draw() {
	win := s.gamedata.Sys.Win
	s.ins.Drawer(win, s.dti.Dt)
	// DEBUG: debug mode
	s.gamedata.Tool.DebugLogger.Draw(win, pixel.IM)
}

// inputHandle : 输入监听
func (s *Scene) inputHandle() {
	s.ins.InputHandle(s.gamedata.Sys.Win, s.dti.Dt)
}
