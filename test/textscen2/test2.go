package textscen2

import (
	"campaign/control/scenario"
	"campaign/control/strdef"
	"fmt"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

const (
	// rps : 执行部分刷新速率
	rps = 2
)

// Scenario : 返回该实例（一般不需要修改）
func Scenario() *scenario.Scenario {
	i := scenario.Instance(&instance{})
	s := scenario.New(rps, i)
	return s
}

// instance : 数据实例，可以自定义
type instance struct {
	sdata *strdef.ShareData
	data  *customData
}

// customData : 自定义数据格式
type customData struct {
	ticker *time.Ticker
}

// SetSData : 设置共享数据指针
func (i *instance) SetSData(sdata *strdef.ShareData) {
	i.sdata = sdata
}

// ResetData : 重置自定义数据
func (i *instance) ResetData() {
	i.data = &customData{}
}

func (i *instance) Initial(w *pixelgl.Window) {
	logger := i.sdata.Tool.DebugLogger
	fmt.Fprintln(logger, "Here is Test 2")
	i.data.ticker = time.NewTicker(2 * time.Second)

}

func (i *instance) Excuter(dt float64) strdef.Request {
	logger := i.sdata.Tool.DebugLogger

	r := strdef.DefaultRequest()
	select {
	case <-i.data.ticker.C:
		r.Continue = false
		r.ResetData = true
		r.NextScenario = "test"
		return r
	default:
	}

	fmt.Fprintf(logger, "%v\n", time.Now().String())

	return r
}

func (i *instance) Drawer(w *pixelgl.Window) {
}
