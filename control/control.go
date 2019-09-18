package control

import (
	"ebby/common/cfgloader"
	"ebby/common/font"
	"ebby/common/logger"
	"ebby/control/def"
	"ebby/errdef"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Control : 控制中心
type Control struct {
	win       *pixelgl.Window
	display   def.DisplayController
	config    *viper.Viper
	logger    *logrus.Logger
	sdata     *def.ShareData
	scenarios ScenarioMap
	now       string
	fn        *Functions
}

// Functions : 外部加载函数
type Functions struct {
	Ini  func(*Control)
	Bfex func(*Control)
}

// New : 返回新的控制中心类
func New(sm ScenarioMap, sd interface{}) *Control {
	config := cfgloader.Init()
	win := setWindow(config)
	c := &Control{
		win:     win,
		display: def.NewDisplayController(win),
		config:  config,
		logger:  logger.New(),
		now:     config.GetString("scenario.entry"),
		fn:      &Functions{},
	}
	c.SetSData(sd)
	c.scenarios = loadScenarios(sm, c.sdata, config)
	return c
}

func setWindow(config *viper.Viper) *pixelgl.Window {

	title := config.GetString("screen.title")
	screenX := config.GetFloat64("screen.rX")
	screenY := config.GetFloat64("screen.rY")
	vSync := config.GetBool("screen.VSync")
	resizable := config.GetBool("screen.resizable")

	cfg := pixelgl.WindowConfig{
		Title:     title,
		Bounds:    pixel.R(0, 0, screenX, screenY),
		Resizable: resizable,
		VSync:     vSync,
	}

	win, err := pixelgl.NewWindow(cfg)
	errdef.CheckErr(err, "control/Enter", errdef.CreateWindow)

	return win
}

// SetSData : 设定 SData
func (c *Control) SetSData(sd interface{}) {
	c.sdata = &def.ShareData{UserData: sd}
	c.sdata.Tool.Logger = c.logger
	c.sdata.Tool.Config = c.config
	c.sdata.Tool.Display = c.display
	if sd == nil {
		c.logger.Warn("[control] 未定义共享数据")
	}
}

// SData : 取得 SData
func (c *Control) SData() *def.ShareData {
	return c.sdata
}

// SetFunctions : 设置外部函数
func (c *Control) SetFunctions(fn *Functions) {
	c.fn = fn
}

// Init : 初始化
func (c *Control) Init() {
	c.fn.Ini(c)
}

// BeforeExit : 关闭前行为（保存数据等）
func (c *Control) BeforeExit() {
	c.fn.Bfex(c)
}

// SetDebugLogger : 使用 debug 模式
func (c *Control) SetDebugLogger() {
	// 加载debug用字符集
	debugAtlas := font.DebugAtlas()
	c.sdata.Resource.DebugAtlas = debugAtlas

	// 加载 debug 用屏幕显示 logger
	locate := pixel.V(4, c.win.Bounds().H()-debugAtlas.LineHeight())
	logger := text.New(locate, debugAtlas)
	c.sdata.Tool.DebugLogger = logger
	c.sdata.Tool.Display.PushShareFn(font.GetDebugLoggerDisplayCallBack(logger))
}

// Run : 运行 scenario
func (c *Control) Run() {
	r := def.DefaultRequest
	for {
		s, ok := c.scenarios[c.now]
		checkScenario(ok, c.now)

		initScenario(s, r, c.win)
		r = s.Run(c.win)

		if r.Terminate {
			return
		}
		c.now = r.NextScenario
	}
}
