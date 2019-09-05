package scenario

import (
	"ebby/control/strdef"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/spf13/viper"
)

// Instance : 场景实例接口，场景数据结构由具体场景定义
type Instance interface {
	Initial(*pixelgl.Window)
	Excuter(float64) strdef.Request
	Drawer(*pixelgl.Window)
	InputHandle(*pixelgl.Window, float64)
	SetSData(*strdef.ShareData)
	ResetData()
}

// Scenario :
type Scenario struct {
	config  *viper.Viper
	rps     int
	sTicker *time.Ticker // sTicker: 总 rps 限制用 ticker，用于节省 cpu
	eTicker *time.Ticker // eTicker: 场景执行部分限制用 ticker
	ins     Instance     // Scenario 实例
	sdata   *strdef.ShareData
}

// New :
func New(rps int, i Instance) *Scenario {
	d := time.Second / time.Duration(rps)
	s := &Scenario{
		rps:     rps,
		eTicker: time.NewTicker(d),
		ins:     i,
	}
	return s
}

// SetConfig : 设置公共配置文件
func (s *Scenario) SetConfig(config *viper.Viper) {
	s.config = config
	// 设置总rps限制用ticker
	rps := config.GetInt("scenario.maxRate")
	d := time.Second / time.Duration(rps)
	s.sTicker = time.NewTicker(d)
}

// Initial : 初始化数据
func (s *Scenario) Initial(w *pixelgl.Window) {
	s.ins.Initial(w)
}

// SetData : 设置并初始化数据
func (s *Scenario) SetData(sdata *strdef.ShareData) {
	s.sdata = sdata
	s.ins.SetSData(sdata)
	s.ResetData()
}

// ResetData : 重置自定义数据
func (s *Scenario) ResetData() {
	s.ins.ResetData()
}
