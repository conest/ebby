package scene

import (
	"github.com/conest/ebby/game/def"

	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

// Run : 主运行
func (s *Scene) Run() def.Request {

	win := s.gamedata.Sys.Win
	debugMode := s.gamedata.Tool.DebugMode

	// 执行循环
	for {
		if win.Closed() {
			return def.Request{Terminate: true}
		}
		s.systemUpdate()

		s.inputHandle()
		if r := s.excute(); !r.Continue {
			s.req = r
			return s.req
		}

		win.Clear(colornames.Black)
		s.draw()
		s.debug(debugMode)
		win.Update()

		<-s.sTicker.C
	}
}

// systemUpdate : 系统数据更新
func (s *Scene) systemUpdate() {
	s.gamedata.Sys.Display.Update()
	s.dti.Update()
}

// excute : 数据执行
func (s *Scene) excute() def.Request {
	return s.ins.Excuter(s.dti.Dt)
}

// draw : 绘图
func (s *Scene) draw() {
	s.ins.Drawer(s.gamedata.Sys.Win, s.dti.Dt)
}

// inputHandle : 输入监听
func (s *Scene) inputHandle() {
	s.ins.InputHandle(s.gamedata.Sys.Win, s.dti.Dt)
}

// debug : debug模式执行部分
func (s *Scene) debug(debugMode bool) {
	if !debugMode {
		return
	}
	s.gamedata.Tool.Fps.Update()
	s.gamedata.Tool.DebugLogger.Draw(s.gamedata.Sys.Win, pixel.IM)
}
