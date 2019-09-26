package scenario

import (
	"ebby/control/def"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/spf13/viper"
)

// Instance : 场景实例接口，场景数据结构由具体场景定义
type Instance interface {
	Initial(*pixelgl.Window)
	Excute(DeltaTime) def.Request
	Drawer(*pixelgl.Window, float64)
	InputHandle(*pixelgl.Window, float64)
	SetSData(*def.ShareData)
	ResetData()
}

// Scenario :
type Scenario struct {
	config  *viper.Viper
	rps     int
	sTicker *time.Ticker // sTicker: 总 rps 限制用 ticker，用于节省 cpu
	eTicker *time.Ticker // eTicker: 场景执行部分限制用 ticker
	ins     Instance     // Scenario 实例
	sdata   *def.ShareData
}

// New :
func New(rps int, instance Instance) *Scenario {
	d := time.Second / time.Duration(rps)
	s := &Scenario{
		rps:     rps,
		eTicker: time.NewTicker(d),
		ins:     instance,
	}
	return s
}

// SetConfig : 设置公共配置文件
func (scenario *Scenario) SetConfig(config *viper.Viper) {
	scenario.config = config
	// 设置总rps限制用ticker
	rps := config.GetInt("scenario.maxRate")
	d := time.Second / time.Duration(rps)
	scenario.sTicker = time.NewTicker(d)
}

// Initial : 初始化数据
func (scenario *Scenario) Initial(w *pixelgl.Window) {
	scenario.ins.Initial(w)
}

// SetData : 设置并初始化数据
func (scenario *Scenario) SetData(sdata *def.ShareData) {
	scenario.sdata = sdata
	scenario.ins.SetSData(sdata)
	scenario.ResetData()
}

// ResetData : 重置自定义数据
func (scenario *Scenario) ResetData() {
	scenario.ins.ResetData()
}

// Timer : Delta Timer，用来生成（每次屏幕刷新之间的）T0 -> T1的时间差
type Timer struct {
	LastTime time.Time
	Delta    float64
}

// NewDeltaTimer : 生成新的 Delta 实例
func NewDeltaTimer() DeltaTimer {
	return DeltaTimer{LastTime: time.Now()}
}

// GetDeltaTime : 获取时间差，并重置LastTime
func (d *DeltaTimer) GetDeltaTime() float64 {
	delta := time.Since(d.LastTime).Seconds()
	d.delta = delta
	d.LastTime = time.Now()
	return delta
}
