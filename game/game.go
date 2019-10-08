package game

import (
	"github.com/conest/ebby/game/def"
	"github.com/conest/ebby/game/sys"
	"github.com/conest/ebby/game/tool"
	"github.com/conest/ebby/model"
	"github.com/conest/ebby/system"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/spf13/viper"
)

// Game : 游戏控制中心
type Game struct {
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
		now: config.GetString("scene.entry"),
		fn:  &ExFunctions{},
	}
	g.SetGameData(win, publicData)
	g.gamedata.Sys.Win = win
	g.gamedata.Sys.Logger = system.NewLogger()
	g.gamedata.Sys.Config = config
	g.gamedata.Sys.Display = sys.NewDisplayController(win)

	if config.GetString("mode") == "debug" {
		g.UseDebugMode()
	}

	g.scenes = loadScenes(sceneMap, g.gamedata)

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
	g.gamedata = &def.GameData{
		Sys:        def.Sys{},
		Tool:       def.Tool{DebugMode: false},
		PublicData: publicData,
	}
	if publicData == nil {
		g.gamedata.Sys.Logger.Warn("[game] 未定义publicData")
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

// UseDebugMode : 使用 debug 模式
func (g *Game) UseDebugMode() {
	g.gamedata.Sys.Logger.Debug("[game] Using debug mode")
	g.gamedata.Tool.DebugMode = true

	// 加载debug用字符集
	debugAtlas := model.DebugAtlas()
	g.gamedata.Tool.DebugAtlas = debugAtlas

	// 加载 debug 用屏幕显示 logger
	locate := pixel.V(4, g.gamedata.Sys.Win.Bounds().H()-debugAtlas.LineHeight())
	logger := text.New(locate, debugAtlas)
	g.gamedata.Tool.DebugLogger = logger
	g.gamedata.Sys.Display.PushPublicFn(model.DebugLoggerDisplayCallBack(logger))

	// 加载 FPS 模块
	g.gamedata.Tool.Fps = tool.NewFps(g.gamedata.Sys.Win, debugAtlas)
}

// terminateScene : 结束场景
func (g *Game) terminateScene() {
	g.gamedata.Sys.Display.ClearSceneFn()
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
