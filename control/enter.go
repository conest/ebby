package control

import (
	"campaign/common/cfgloader"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Enter : 创建窗口并初始化控制中心
func Enter() {

	config := cfgloader.Init()

	title := config.GetString("screen.title")
	screenX := config.GetFloat64("screen.rX")
	screenY := config.GetFloat64("screen.rY")
	vSync := config.GetBool("screen.VSync")

	cfg := pixelgl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(0, 0, screenX, screenY),
		VSync:  vSync,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// win.SetCursorVisible(false)

	// 创建新的控制中心
	c := New(win, config)
	c.DebugMode()
	c.Init()
	c.Run()
	c.BeforeExit()
}
