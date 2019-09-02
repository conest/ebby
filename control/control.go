package control

import (
	"campaign/common/font"
	"campaign/common/logger"
	"campaign/control/scenario"
	"campaign/control/strdef"
	"campaign/errdef"
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/spf13/viper"
)

const (
	// DEBUG: test -> start
	iniScenario = "test"
)

// Control : 控制中心
type Control struct {
	win       *pixelgl.Window
	config    *viper.Viper
	sdata     *strdef.ShareData
	scenarios ScenarioListMap
	now       string
}

// New : 返回新的控制中心类
func New(win *pixelgl.Window, config *viper.Viper) *Control {
	sdata := &strdef.ShareData{}
	c := &Control{
		win:       win,
		config:    config,
		sdata:     sdata,
		scenarios: ScenarioList(sdata, config),
		now:       iniScenario,
	}
	return c
}

// Init : 初始化
func (c *Control) Init() {
	c.sdata.Tool.Logger = logger.New()
	c.sdata.Tool.Logger.Info("[Control Center] Started")
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
	r := strdef.DefaultRequest()
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

// BeforeExit : 关闭前行为（保存数据等）
func (c *Control) BeforeExit() {
	c.sdata.Tool.Logger.Info("[control] Ternimated")
}

func initScenario(s *scenario.Scenario, r strdef.Request, win *pixelgl.Window) {
	if r.ResetData {
		s.ResetData()
		s.Initial(win)
	}
}

// checkScenario : 检测对应的scenario是否存在
func checkScenario(ok bool, id string) {
	if !ok {
		log := logger.New()
		errString := fmt.Sprintf("[control]<%s> %v", errdef.NoScenario.Str, id)
		log.Error(errString)
		panic(errString)
	}
}
