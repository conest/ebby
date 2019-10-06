package game

import (
	"github.com/conest/ebby/game/def"
	"github.com/conest/ebby/game/sys"
	"github.com/conest/ebby/model"
	"github.com/conest/ebby/system"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Game : 游戏控制中心
type Game struct {
	win      *pixelgl.Window
	display  sys.DisplayController
	config   *viper.Viper
	logger   *logrus.Logger
	gamedata *def.GameData
	scenes   SceneMap
	now      string
	fn       *ExFunctions
}

// ExFunctions : 外部加载函数
type ExFunctions struct {
	Ini func(*Game)
	Exi func(*Game)
}

// New : 返回新的Game实例
func New(sceneMap SceneMap, publicData interface{}) *Game {
	config := system.ViperInit()
	win := createWindow(config)
	g := &Game{
		win:     win,
		display: sys.NewDisplayController(win),
		config:  config,
		logger:  system.NewLogger(),
		now:     config.GetString("scene.entry"),
		fn:      &ExFunctions{},
	}
	g.SetGameData(win, publicData)
	g.scenes = loadScenes(sceneMap, g.gamedata, config)
	return g
}

func createWindow(config *viper.Viper) *pixelgl.Window {

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

// SetGameData : 设定 GameData
func (g *Game) SetGameData(win *pixelgl.Window, publicData interface{}) {
	g.gamedata = &def.GameData{PublicData: publicData}
	g.gamedata.Sys.Win = win
	g.gamedata.Sys.Logger = g.logger
	g.gamedata.Sys.Config = g.config
	g.gamedata.Sys.Display = g.display
	if publicData == nil {
		g.logger.Warn("[game] 未定义publicData")
	}
}

// GameData : 取得 GameData
func (g *Game) GameData() *def.GameData {
	return g.gamedata
}

// SetFunctions : 设置外部函数
func (g *Game) SetFunctions(fn *ExFunctions) {
	g.fn = fn
}

// Init : 初始化
func (g *Game) Init() {
	g.fn.Ini(g)
}

// BeforeExit : 关闭前行为（保存数据等）
func (g *Game) BeforeExit() {
	g.fn.Exi(g)
}

// SetDebugLogger : 使用 debug 模式
func (g *Game) SetDebugLogger() {
	// 加载debug用字符集
	debugAtlas := model.DebugAtlas()
	g.gamedata.Tool.DebugAtlas = debugAtlas

	// 加载 debug 用屏幕显示 logger
	locate := pixel.V(4, g.win.Bounds().H()-debugAtlas.LineHeight())
	logger := text.New(locate, debugAtlas)
	g.gamedata.Tool.DebugLogger = logger
	g.gamedata.Sys.Display.PushPublicFn(model.GetDebugLoggerDisplayCallBack(logger))
}

// terminateScene : 结束场景
func (g *Game) terminateScene() {
	g.display.ClearSceneFn()
}

// Run : 运行 scene
func (g *Game) Run() {
	r := def.DefaultRequest
	for {
		scene, ok := g.scenes[g.now]
		checkScene(ok, g.now)

		initScene(scene, r)
		r = scene.Run()

		if r.Terminate {
			return
		}
		g.terminateScene()
		g.now = r.NextScene
	}
}
