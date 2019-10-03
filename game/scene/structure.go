package scene

import (
	"github.com/conest/ebby/game/def"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/spf13/viper"
)

// Instance : 场景实例接口，场景数据结构由具体场景定义
type Instance interface {
	Initial(*pixelgl.Window)
	Excuter(DeltaTime) def.Request
	Drawer(*pixelgl.Window, float64)
	InputHandle(*pixelgl.Window, float64)
	SetSData(*def.ShareData)
	ResetData()
}

// Scene :
type Scene struct {
	config  *viper.Viper
	rps     int
	sTicker *time.Ticker // sTicker: 总 rps 限制用 ticker，用于节省 cpu
	eTicker *time.Ticker // eTicker: 场景执行部分限制用 ticker
	ins     Instance     // Scene 实例
	sdata   *def.ShareData
}

// New :
func New(rps int, i Instance) *Scene {
	d := time.Second / time.Duration(rps)
	s := &Scene{
		rps:     rps,
		eTicker: time.NewTicker(d),
		ins:     i,
	}
	return s
}

// SetConfig : 设置公共配置文件
func (s *Scene) SetConfig(config *viper.Viper) {
	s.config = config
	// 设置总rps限制用ticker
	rps := config.GetInt("scene.maxRate")
	d := time.Second / time.Duration(rps)
	s.sTicker = time.NewTicker(d)
}

// Initial : 初始化数据
func (s *Scene) Initial(w *pixelgl.Window) {
	s.ins.Initial(w)
}

// SetData : 设置并初始化数据
func (s *Scene) SetData(sdata *def.ShareData) {
	s.sdata = sdata
	s.ins.SetSData(sdata)
	s.ResetData()
}

// ResetData : 重置自定义数据
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

// Get : 重置自定义数据
func (d *DeltaTime) Get() float64 {
	dt := time.Since(d.Last).Seconds()
	d.Dt = dt
	d.Last = time.Now()
	return dt
}
