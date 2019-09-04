package control

import (
	"ebby/common/cfgloader"
	"ebby/errdef"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/spf13/viper"
)

// Enter : 创建窗口并初始化控制中心
func Enter(sm ScenarioMap, fn *Functions) {

	config, win := setWindow()

	// 创建新的控制中心
	c := New(win, config, sm)
	c.SetFunctions(fn)
	c.DebugMode()
	c.Init()
	c.Run()
	c.BeforeExit()
}

func setWindow() (*viper.Viper, *pixelgl.Window) {
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
	errdef.CheckErr(err, "control/Enter", errdef.CreateWindow)

	return config, win
}
