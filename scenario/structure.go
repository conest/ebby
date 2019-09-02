package scenario

import (
	"campaign/control/datastruct"
	"campaign/control/order"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/spf13/viper"
)

// Instance : 场景实例接口，场景数据结构由具体场景定义
type Instance interface {
	Initial(*pixelgl.Window)
	Excuter(float64) *order.Request
	Drawer(*pixelgl.Window)
	SetSData(*datastruct.ShareData)
}

// Scenario :
type Scenario struct {
	config  *viper.Viper
	rps     int
	sTicker *time.Ticker // sTicker: 总 rps 限制用 ticker，用于节省 cpu
	eTicker *time.Ticker // eTicker: 场景执行部分限制用 ticker
	ins     Instance     // Scenario 实例
	sdata   *datastruct.ShareData
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
	rps := config.GetInt("scenario.maxRate")
	d := time.Second / time.Duration(rps)
	s.sTicker = time.NewTicker(d)
	s.config = config
}

// SetSData : 设置共享数据
func (s *Scenario) SetSData(sdata *datastruct.ShareData) {
	s.sdata = sdata
	s.ins.SetSData(sdata)
}
