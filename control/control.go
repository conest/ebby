package control

import (
	"ebby/common/font"
	"ebby/common/logger"
	"ebby/control/def"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Control : 控制中心
type Control struct {
	win       *pixelgl.Window
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
func New(win *pixelgl.Window, config *viper.Viper, sm ScenarioMap, sd interface{}) *Control {
	c := &Control{
		win:    win,
		config: config,
		logger: logger.New(),
		now:    config.GetString("scenario.entry"),
		fn:     &Functions{},
	}
	c.sdata = setSData(sd, c.logger, config)
	c.scenarios = loadScenarios(sm, c.sdata, config)
	return c
}

func setSData(sd interface{}, logger *logrus.Logger, config *viper.Viper) *def.ShareData {
	sdata := &def.ShareData{UserData: sd}
	sdata.Tool.Logger = logger
	sdata.Tool.Config = config
	if sd == nil {
		logger.Warn("[control] 未定义共享数据")
	}
	return sdata
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

// DebugMode : 使用 debug 模式
func (c *Control) DebugMode() {
	// 加载debug用字符集
	debugAtlas := font.GetDebugAtlas()
	c.sdata.Resource.DebugAtlas = debugAtlas

	// 加载 debug 用屏幕显示 logger
	locate := pixel.V(4, c.win.Bounds().H()-debugAtlas.LineHeight())
	c.sdata.Tool.DebugLogger = text.New(locate, debugAtlas)
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
