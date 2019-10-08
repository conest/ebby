package scene

import (
	"time"

	"github.com/conest/ebby/game/def"

	"github.com/faiface/pixel/pixelgl"
)

// Instance : 场景实例接口，场景数据结构由具体场景定义
type Instance interface {
	Initial(*pixelgl.Window)
	Excuter(float64) def.Request
	Drawer(*pixelgl.Window, float64)
	InputHandle(*pixelgl.Window, float64)
	SetGameData(*def.GameData)
	ResetData()
}

// Scene :
type Scene struct {
	req      def.Request
	dti      DeltaTime
	sTicker  *time.Ticker // sTicker: 总 rps 限制用 ticker，用于节省 cpu
	ins      Instance     // Scene 实例
	gamedata *def.GameData
}

// New :
func New(i Instance) *Scene {
	return &Scene{
		req: def.DefaultRequest,
		dti: NewDT(),
		ins: i,
	}
}

// SetGameData : 设置 gamedata
func (s *Scene) SetGameData(gamedata *def.GameData) {
	s.gamedata = gamedata
	s.ins.SetGameData(gamedata)
	s.makeTicker()
}

// InitInstance : 初始化场景 Instance
func (s *Scene) InitInstance() {
	s.ins.Initial(s.gamedata.Sys.Win)
}

// ResetInstanceData : 重置Instance数据
func (s *Scene) ResetInstanceData() {
	s.ins.ResetData()
}

// makeTicker : 配置循环限制用 Ticker
func (s *Scene) makeTicker() {
	// 设置总rps限制用ticker
	rps := s.gamedata.Sys.Config.GetInt("scene.maxRate")
	d := time.Second / time.Duration(rps)
	s.sTicker = time.NewTicker(d)
}
