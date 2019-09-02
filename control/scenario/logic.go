package scenario

import (
	"campaign/control/strdef"
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

// Run : 主运行
func (s *Scenario) Run(w *pixelgl.Window) strdef.Request {

	req := strdef.DefaultRequest()
	last := time.Now()
	// DEBUG: debug mode
	// fps
	var (
		frames    = 0
		fpsTicker = time.NewTicker(time.Second)
		fpsTxt    = text.New(pixel.V(4, 4), s.sdata.Resource.DebugAtlas)
	)

	// 执行循环
	for {
		if w.Closed() {
			req = strdef.Request{Terminate: true}
			return req
		}

		// Delta Time
		dt := time.Since(last).Seconds()
		last = time.Now()
		if r := s.excute(dt); !r.Continue {
			req = r
			return req
		}

		w.Clear(colornames.Black)
		s.draw(w)

		// DEBUG: debug mode
		fps(w, fpsTxt, &frames, fpsTicker.C)
		fpsTxt.Draw(w, pixel.IM)

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

// DEBUG: debug mode
func fps(win *pixelgl.Window, fpsTxt *text.Text, frames *int, tick <-chan time.Time) {
	*frames++
	select {
	case <-tick:
		fpsTxt.Clear()
		fmt.Fprintf(fpsTxt, "FPS: %v", *frames)
		*frames = 0
	default:
	}
}
