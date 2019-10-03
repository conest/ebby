package game

import (
	"ebby/game/def"
	"ebby/game/sys"
	"ebby/model"
	"ebby/system"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Game : 游戏控制中心
type Game struct {
	win     *pixelgl.Window
	display sys.DisplayController
	config  *viper.Viper
	logger  *logrus.Logger
	sdata   *def.ShareData
	scenes  SceneMap
	now     string
	fn      *ExFunctions
}

// ExFunctions : 外部加载函数
type ExFunctions struct {
	Ini func(*Game)
	Exi func(*Game)
}

// New : 返回新的控制中心类
func New(sm SceneMap, sd interface{}) *Game {
	config := system.ViperInit()
	win := setWindow(config)
	c := &Game{
		win:     win,
		display: sys.NewDisplayController(win),
		config:  config,
		logger:  system.NewLogger(),
		now:     config.GetString("scene.entry"),
		fn:      &ExFunctions{},
	}
	c.SetSData(sd)
	c.scenes = loadScenes(sm, c.sdata, config)
	return c
}

func setWindow(config *viper.Viper) *pixelgl.Window {

	title := config.GetString("screen.title")
	screenX := config.GetFloat64("screen.rX")
	screenY := config.GetFloat64("screen.rY")
	vSync := config.GetBool("screen.vSync")
	resizable := config.GetBool("screen.resizable")

	cfg := pixelgl.WindowConfig{
		Title:     title,
		Bounds:    pixel.R(0, 0, screenX, screenY),
		Resizable: resizable,
		VSync:     vSync,
	}

	win, err := pixelgl.NewWindow(cfg)
	system.CheckErr(err, "game/Enter", system.ErrorTable["CreateWindow"])

	return win
}

// SetSData : 设定 SData
func (c *Game) SetSData(sd interface{}) {
	c.sdata = &def.ShareData{UserData: sd}
	c.sdata.Tool.Logger = c.logger
	c.sdata.Tool.Config = c.config
	c.sdata.Tool.Display = c.display
	if sd == nil {
		c.logger.Warn("[game] 未定义共享数据")
	}
}

// SData : 取得 SData
func (c *Game) SData() *def.ShareData {
	return c.sdata
}

// SetFunctions : 设置外部函数
func (c *Game) SetFunctions(fn *ExFunctions) {
	c.fn = fn
}

// Init : 初始化
func (c *Game) Init() {
	c.fn.Ini(c)
}

// BeforeExit : 关闭前行为（保存数据等）
func (c *Game) BeforeExit() {
	c.fn.Exi(c)
}

// SetDebugLogger : 使用 debug 模式
func (c *Game) SetDebugLogger() {
	// 加载debug用字符集
	debugAtlas := model.DebugAtlas()
	c.sdata.Resource.DebugAtlas = debugAtlas

	// 加载 debug 用屏幕显示 logger
	locate := pixel.V(4, c.win.Bounds().H()-debugAtlas.LineHeight())
	logger := text.New(locate, debugAtlas)
	c.sdata.Tool.DebugLogger = logger
	c.sdata.Tool.Display.PushShareFn(model.GetDebugLoggerDisplayCallBack(logger))
}

// terminateScene : 结束场景
func (c *Game) terminateScene() {
	c.display.ClearSceneFn()
}

// Run : 运行 scene
func (c *Game) Run() {
	r := def.DefaultRequest
	for {
		s, ok := c.scenes[c.now]
		checkScene(ok, c.now)

		initScene(s, r, c.win)
		r = s.Run(c.win)

		if r.Terminate {
			return
		}
		c.terminateScene()
		c.now = r.NextScene
	}
}
