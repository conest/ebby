package strdef

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

// Fps : FPS显示组建
type Fps struct {
	win    *pixelgl.Window
	frames int
	ticker *time.Ticker
	txt    *text.Text
}

// NewFps : 创建新的FPS
func NewFps(w *pixelgl.Window, atlas *text.Atlas) Fps {
	return Fps{
		win:    w,
		frames: 0,
		ticker: time.NewTicker(time.Second),
		txt:    text.New(pixel.V(4, 4), atlas),
	}
}

// Update : 更新FPS状态
func (f *Fps) Update() {
	f.frames++
	select {
	case <-f.ticker.C:
		f.txt.Clear()
		fmt.Fprintf(f.txt, "FPS: %v", f.frames)
		f.frames = 0
	default:
	}
	f.txt.Draw(f.win, pixel.IM)
}
