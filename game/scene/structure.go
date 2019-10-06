package scene

import (
	"time"

	"github.com/conest/ebby/game/def"

	"github.com/faiface/pixel/pixelgl"
	"github.com/spf13/viper"
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
	config   *viper.Viper
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

// SetConfig : 设置公共配置文件
func (s *Scene) SetConfig(config *viper.Viper) {
	s.config = config
	// 设置总rps限制用ticker
	rps := config.GetInt("scene.maxRate")
	d := time.Second / time.Duration(rps)
	s.sTicker = time.NewTicker(d)
}

// SetData : 设置并初始化数据
func (s *Scene) SetData(gamedata *def.GameData) {
	s.gamedata = gamedata
	s.ins.SetGameData(gamedata)
	s.ResetData()
}

// IniInstance : 初始化场景 Instance
func (s *Scene) IniInstance() {
	s.ins.Initial(s.gamedata.Sys.Win)
}

// ResetData : 重置Instance数据
func (s *Scene) ResetData() {
	s.ins.ResetData()
}

// DeltaTime : Delta Time，每次屏幕刷新之间的时间差
type DeltaTime struct {
	Last time.Time
	Dt   float64
}

// NewDT : 生成新的 Delta 实例
func NewDT() DeltaTime {
	return DeltaTime{Last: time.Now()}
}

// Update : 刷新并返回Delta Time
func (d *DeltaTime) Update() float64 {
	dt := time.Since(d.Last).Seconds()
	d.Dt = dt
	d.Last = time.Now()
	return dt
}
