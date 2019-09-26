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

// CreateControl : 返回新的控制中心类
func CreateControl(scenarioMap ScenarioMap, sData interface{}) *Control {
	config := cfgloader.Init()
	win := setWindow(config)
	ctrl := &Control{
		win:     win,
		display: def.NewDisplayController(win),
		config:  config,
		logger:  logger.New(),
		now:     config.GetString("scenario.entry"),
		fn:      &Functions{},
	}
	ctrl.SetSData(sData)
	ctrl.scenarios = loadScenarios(scenarioMap, ctrl.sdata, config)
	return ctrl
}

// SetUpWindow : Set up a window(canvas?) by viper
func SetUpWindow(config *viper.Viper) *pixelgl.Window {
	title := config.GetString("screen.title")
	screenX := config.GetFloat64("screen.rX")
	screenY := config.GetFloat64("screen.rY")
	vSync := config.GetBool("screen.VSync")
	resizable := config.GetBool("screen.resizable")

	windowConfig := pixelgl.WindowConfig{
		Title:     title,
		Bounds:    pixel.R(0, 0, screenX, screenY),
		Resizable: resizable,
		VSync:     vSync,
	}

	window, err := pixelgl.NewWindow(windowConfig)
	errdef.CheckErr(err, "control/Enter", errdef.CreateWindow)

	return window
}

// SetSData : 设定 SData
func (ctrl *Control) SetSData(sd interface{}) {
	ctrl.sdata = &def.ShareData{UserData: sd}
	ctrl.sdata.Tool.Logger = ctrl.logger
	ctrl.sdata.Tool.Config = ctrl.config
	ctrl.sdata.Tool.Display = ctrl.display
	if sd == nil {
		ctrl.logger.Warn("[control] 未定义共享数据")
	}
}

// GetSData : 取得 SData
func (ctrl *Control) GetSData() *def.ShareData {
	return ctrl.sdata
}

// SetFunctions : 设置外部函数
func (ctrl *Control) SetFunctions(fn *Functions) {
	ctrl.fn = fn
}

// Init : 初始化
func (ctrl *Control) Init() {
	ctrl.fn.Ini(c)
}

// BeforeExit : 关闭前行为（保存数据等）
func (ctrl *Control) BeforeExit() {
	ctrl.fn.Bfex(c)
}

// SetDebugLogger : 使用 debug 模式
func (ctrl *Control) SetDebugLogger() {
	// 加载debug用字符集
	debugAtlas := font.DebugAtlas()
	ctrl.sdata.Resource.DebugAtlas = debugAtlas

	// 加载 debug 用屏幕显示 logger
	locate := pixel.V(4, ctrl.win.Bounds().H()-debugAtlas.LineHeight())
	logger := text.New(locate, debugAtlas)
	ctrl.sdata.Tool.DebugLogger = logger
	ctrl.sdata.Tool.Display.PushShareFn(font.GetDebugLoggerDisplayCallBack(logger))
}

// Run : 运行 scenario
func (ctrl *Control) Run() {
	req := def.DefaultRequest
	for {
		scenario, ok := c.scenarios[c.now]
		checkScenario(ok, ctrl.now)

		initScenario(scenario, req, ctrl.win)
		req = s.Run(c.win)

		if req.Terminate {
			return
		}
		ctrl.now = req.NextScenario
	}
}
